package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Отобразите страницу авторизации с кнопками выбора социальной сети
	// Например:
	http.ServeFile(w, r, "static/index.html")
}

func vkLoginHandler(w http.ResponseWriter, r *http.Request) {
	conf := GetConfig()
	vkAuthURL := fmt.Sprintf("https://oauth.vk.com/authorize?client_id=%s&display=page&redirect_uri=%s&response_type=code&v=5.74", conf.VK_APP_ID, conf.VK_REDIRECT_URL)
	http.Redirect(w, r, vkAuthURL, http.StatusFound)
}

func vkCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Код авторизации не найден", http.StatusBadRequest)
		return
	}
	conf := GetConfig()
	// Запрос на получение токена ВКонтакте
	tokenURL := fmt.Sprintf(
		"https://oauth.vk.com/access_token?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s",
		conf.VK_APP_ID, conf.VK_SECRET_KEY, conf.VK_REDIRECT_URL, code)
	resp, err := http.Get(tokenURL)
	if err != nil {
		http.Error(w, "Ошибка при запросе токена у ВКонтакте", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения ответа от ВКонтакте", http.StatusInternalServerError)
		return
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON от ВКонтакте", http.StatusInternalServerError)
		return
	}

	// Здесь можно сохранить полученный access token в базу данных или использовать его для запросов к API ВКонтакте от имени пользователя
	// Например:
	fmt.Fprintf(w, "Токен ВКонтакте: %s", tokenResp.AccessToken)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/callback/vk", vkCallbackHandler)
	http.HandleFunc("/login/vk", vkLoginHandler)

	http.ListenAndServe(":8080", nil)
}
