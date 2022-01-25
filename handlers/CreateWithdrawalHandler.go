package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	"net/http"
	"strconv"
)

func CreateWithdrawalHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("CreateWithdrawalHandler")

	MakeSureLoggedIn(response, request)

	data := struct {
		Errors []string
	}{
		Errors: []string{},
	}

	if request.Method == "POST" {
		CreateWithdrawalPostHandler(response, request, &data)
	}

	render, err := template.ParseFiles("templates/create-withdrawal.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}

//Functie om te checken of er genoeg geld is om te withdrawen (CanWithDrawFromAccount). Als er genoeg is, wordt het gelijk afgeschreven en doorgestuurd (WithdrawFromAccount)
func CreateWithdrawalPostHandler(response http.ResponseWriter, request *http.Request, data *struct{ Errors []string }) {
	request.ParseForm()

	id := helpers.GetAccountIdCookie(request)
	amount, _ := strconv.ParseFloat(request.FormValue("amount"), 64)
	description := (request.FormValue("description"))

	var canwithdraw bool = repositories.CanWithdrawFromAccount(id, amount)

	if canwithdraw {
		repositories.CanWithdrawFromAccount(id, amount)
		repositories.WithdrawFromAccount(id, amount, description)
		log.Println("Geld afgeschreven (van, hoeveelheid): ", id, amount)

		http.Redirect(response, request, "/withdrawal-confirmed", http.StatusFound)
	} else {
		log.Println("Kon geen geld afschrijven")
		data.Errors = append(data.Errors, "Onvoldoende balans")
	}
}
