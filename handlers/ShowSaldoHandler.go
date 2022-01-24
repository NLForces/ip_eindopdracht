package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	"net/http"
)

func ShowSaldoHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("ShowSaldoHandler")

	MakeSureLoggedIn(response, request)

	var saldo float64 = repositories.GetSaldoForAccount(helpers.GetAccountIdCookie(request))

	data := struct {
		Saldo float64
	}{
		Saldo: saldo,
	}

	render, err := template.ParseFiles("templates/show-saldo.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}
