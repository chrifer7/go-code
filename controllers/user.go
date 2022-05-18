package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

//Implicit implement of Hendler interface (Golang feature) using the same name and signature
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

func newUserController() *userController {
	//constructing the object (struct) using only memory address
	return &userController{
		//the path should be localhost:3000/users/:id
		//FIX: Go: Getting Started > 6 Creating Functions and Methods > Demo: Starting the Webservice
		userIDPattern: regexp.MustCompile("^/users/(\\d+)/?"),
	}
}
