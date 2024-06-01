package controllers

import (
	"fmt"
	"net/http"
)

func AddPostFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddPostFunc")
}