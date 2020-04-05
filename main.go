package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jain-chetan/users-service/interfaces"
	"github.com/jain-chetan/users-service/lib/database"
	"github.com/jain-chetan/users-service/model"
	api "github.com/jain-chetan/users-service/receiver"
)

func main() {
	err := initDBClient()
	if err != nil {
		log.Fatal("DB Error ", err)
	}

	//router initialization
	router := mux.NewRouter()

	router.HandleFunc("/users", api.Post.PostUserHandler).Methods("POST")
	router.HandleFunc("/users", api.Get.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", api.Put.PutUserHandler).Methods("PUT")
	router.HandleFunc("/users", api.Delete.DeleteUserHandler).Methods("DELETE")
	log.Println("Starting application")
	http.ListenAndServe(":8003", router)
}

func initDBClient() error {
	var config model.DBConfig

	//Read DB credentials from environment variables
	config.User = os.Getenv("DBUSER")
	config.Port = os.Getenv("PORT")
	config.Host = os.Getenv("HOST")

	interfaces.DBClient = new(database.DBRepo)
	err := interfaces.DBClient.DBConnect(config)
	return err
}
