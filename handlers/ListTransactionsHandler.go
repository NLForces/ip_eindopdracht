package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	"net/http"
)

//Vult transacties variabele met de Transactions struct en het ingelogde account ID.
func ListTransactionsHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("ListTransactionsHandler")

	MakeSureLoggedIn(response, request)

	transacties := repositories.GetTransactionsForAccount(helpers.GetAccountIdCookie(request))

	render, err := template.ParseFiles("templates/list-transactions.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, transacties)
}
