package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-contacts/app"
	"go-contacts/controllers"
	"go-contacts/models"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/posts", models.GetPosts).Methods("GET")
	router.HandleFunc("/post/{id}", models.GetPost).Methods("GET")
	router.HandleFunc("/users", models.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", models.GetUserr).Methods("GET")

	router.HandleFunc("/create/user", models.CreateUser).Methods("POST")
	router.HandleFunc("/create/post", models.CreatePost).Methods("POST")

	router.HandleFunc("/delete/user/{id}", models.DeleteUser).Methods("DELETE")
	router.HandleFunc("/delete/post/{id}", models.DeletePost).Methods("DELETE")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
