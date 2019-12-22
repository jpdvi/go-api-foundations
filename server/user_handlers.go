package server

import (
	"encoding/json"
	"fmt"
	"github.com/jpdvi/chat-service/common"
	"github.com/jpdvi/chat-service/models"
	"io/ioutil"
	"net/http"
)

func HandleCreateUser(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonBody, jsonErr := ioutil.ReadAll(r.Body)
		if jsonErr != nil {
			fmt.Println("Error Reading Json")
		}
		var newUser models.User
		err := json.Unmarshal(jsonBody, &newUser)
		if err != nil {
			fmt.Println("Error Parsing Json")
		}
		fmt.Println(newUser.Username)
		env.DB.Where(&models.User{Username: newUser.Username}).First(&newUser)
		if newUser.ID == 0 {
			passwordHash, err := common.CreatePasswordHash(newUser.Password)
			if err != nil {
				fmt.Println("error")
			}
			newUser.Password = passwordHash
			dbObj := env.DB.Create(&newUser)
			json.NewEncoder(w).Encode(dbObj)
		} else {
			http.Error(w, "User Exists", 500)
		}
	})
}

func handleGetUser(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

	})
}
