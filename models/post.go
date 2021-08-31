package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/gorilla/mux"
)

type Post struct {
	gorm.Model

	Title      string
	Body       string
	UserID     int
}

var err error

var GetPost = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post Post

	GetDB().First(&post, params["id"])

	json.NewEncoder(w).Encode(&post)
}

var GetPosts = func(w http.ResponseWriter, r *http.Request) {
	var posts []Post

	GetDB().Find(&posts)

	json.NewEncoder(w).Encode(&posts)
}

var CreatePost = func(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	createdPost := GetDB().Create(&post)
	err = createdPost.Error
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(&createdPost)
}

var DeletePost = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var post Post

	GetDB().First(&post, params["id"])
	GetDB().Delete(&post)

	json.NewEncoder(w).Encode(&post)
}
