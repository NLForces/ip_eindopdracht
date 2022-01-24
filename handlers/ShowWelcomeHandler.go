package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	. "main/types"
	"net/http"
)

func ShowWelcomeHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("ShowWelcomeHandler")

	MakeSureLoggedIn(response, request)

	account := repositories.GetAccount(helpers.GetAccountIdCookie(request))

	data := struct {
		Account Account
	}{
		Account: account,
	}

	render, err := template.ParseFiles("templates/show-welcome.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}