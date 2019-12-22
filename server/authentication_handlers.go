package server

import (
	"encoding/json"
	"github.com/jpdvi/exch-desktop/common"
	"github.com/jpdvi/exch-desktop/models"
	"io/ioutil"
	"net/http"
)

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleLogin(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonBody, jsonErr := ioutil.ReadAll(r.Body)
		if jsonErr != nil {
			http.Error(w, "Malformed Request", 500)
		}
		var login LoginModel
		marshalErr := json.Unmarshal(jsonBody, &login)
		if marshalErr != nil {
			http.Error(w, "Invalid Fields", 500)
		}
		var userModel models.User
		env.DB.Where(&models.User{Username: login.Username}).First(&userModel)
		passwordCorrect, err := common.VerifyPassword(userModel.Password, login.Password)
		if err == nil && passwordCorrect {
			token := common.Token{AccessToken: "alksdjf", RefreshToken: "kalsjdf"}
			js, err := json.Marshal(token)
			if err != nil {
				http.Error(w, "Error Creating Token", 500)
			}
			w.Write(js)
		} else {
			http.Error(w, "Invalid username or password", 401)
		}
	})
}
