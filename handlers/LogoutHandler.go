package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"net/http"
)

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Logged out")
	helpers.SetAccountIdCookie(response, "0")

	render, err := template.ParseFiles("templates/logout.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, nil)
}
