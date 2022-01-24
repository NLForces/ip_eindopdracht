package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func WithdrawalConfirmedHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("WithdrawalConfirmedHandler")

	MakeSureLoggedIn(response, request)

	data := struct{}{}

	render, err := template.ParseFiles("templates/withdrawal-confirmed.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}
