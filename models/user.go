package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/gorilla/mux"

)

type User struct {
	gorm.Model

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Posts []Post
}
var errU error

var GetUserr = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	var posts []Post

	GetDB().First(&user, params["id"])
	GetDB().Model(&user).Related(&posts)

	user.Posts = posts

	json.NewEncoder(w).Encode(&user)
}

var GetUsers = func(w http.ResponseWriter, r *http.Request) {
	var user []User

	GetDB().Find(&user)

	json.NewEncoder(w).Encode(&user)
}

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := GetDB().Create(&user)
	errU = createdUser.Error
	if errU != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(&createdUser)
}

var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user User

	GetDB().First(&user, params["id"])
	GetDB().Delete(&user)

	json.NewEncoder(w).Encode(&user)
}


