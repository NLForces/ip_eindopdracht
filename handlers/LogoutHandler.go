package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Logged out")
	http.SetCookie(response, nil)

	render, err := template.ParseFiles("templates/list-transactions.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, nil)
}
