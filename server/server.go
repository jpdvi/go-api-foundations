package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"os"
)

type Env struct {
	DB *gorm.DB
	R  *mux.Router
}

func CreateEnv() (*Env, error) {
	db, err := gorm.Open("postgres",
		fmt.Sprint(
			" host=", os.Getenv("DBHOST"),
			" port=", os.Getenv("DBPORT"),
			" user=", os.Getenv("DBUSER"),
			" dbname=", os.Getenv("DBNAME"),
			" password=", os.Getenv("DBPASSWORD"),
			" sslmode=disable"))

	r := mux.NewRouter()
	env := &Env{DB: db, R: r}

	if err != nil || env.DB == nil {
		return env, err
	}
	return env, nil
}

func Server() {
	env, err := CreateEnv()
	if err != nil {
		log.Fatal("Error Initializing Environment")
	}
	log.Println("Running Server on port : ", os.Getenv("APPLICATIONPORT"))
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/ws", handleWsConnection)
	env.R.Handle("/user", HandleCreateUser(env)).Methods("POST")
	env.R.Handle("/login", handleLogin(env)).Methods("POST")
	go handleMessages()
	http.ListenAndServe(fmt.Sprint(":", os.Getenv("APPLICATIONPORT")), env.R)
}
