package handlers

import (
	"html/template"
	"log"
	"main/repositories"
	"net/http"
)

func CreationConfirmedHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("CreationConfirmed")

	allAccounts := repositories.GetAllAccounts()
	account_id := allAccounts[len(allAccounts)-1].ID //Roept het ID op van het laatste account en convert dit in een string
	// name := allAccounts[len(allAccounts)-1].Name
	// code := allAccounts[len(allAccounts)-1].Code
	// maxcredit := allAccounts[len(allAccounts)-1].Maxcredit

	data := struct {
		ID uint
	}{
		ID: account_id,
	}

	render, err := template.ParseFiles("templates/account-created.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}
