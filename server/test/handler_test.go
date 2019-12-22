package server

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"github.com/jpdvi/chat-service/models"
	"github.com/jpdvi/chat-service/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	env, err := server.CreateEnv()
	if err != nil {
		t.Fatal(err)
	}
	name := make([]byte, 8)
	rand.Read(name)
	newUser := models.User{Username: string(name), Password: "bigpassword"}
	jsonBody, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := server.HandleCreateUser(env)
	handler.ServeHTTP(rr, req)
	rBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	if status := rr.Code; status == http.StatusInternalServerError {
		t.Errorf("Handler Returned %v Error want %v : Error Message : %v", rr.Code, http.StatusOK, string(rBody))
	}
}
