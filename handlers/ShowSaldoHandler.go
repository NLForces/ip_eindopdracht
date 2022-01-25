package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	"net/http"
)

//Functie om het saldo te laten zien. Vraagt welk account is ingelogd, en draait GetSaldoForAccount met dit account ID. De return wordt geparsed naar show-saldo.
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
