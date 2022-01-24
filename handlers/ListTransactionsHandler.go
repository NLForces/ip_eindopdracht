package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	"net/http"
)

func ListTransactionsHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("ListTransactionsHandler")

	MakeSureLoggedIn(response, request)

	// Eindopdracht 2. Haal hier de lijst van transactions op en geef deze mee aan de template engine

	transacties := repositories.GetTransactionsForAccount(helpers.GetAccountIdCookie(request))

	render, err := template.ParseFiles("templates/list-transactions.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, transacties)
}
