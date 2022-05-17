package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	//without parameters
	http.Handle("/users", *uc)
	//plus user id param
	http.Handle("/users/", *uc)
}
