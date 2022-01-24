package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func TransferConfirmedHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("TransferConfirmedHandler")

	MakeSureLoggedIn(response, request)

	data := struct{}{}

	render, err := template.ParseFiles("templates/transfer-confirmed.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}
