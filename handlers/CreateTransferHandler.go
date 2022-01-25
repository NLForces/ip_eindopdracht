package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	"net/http"
	"strconv"
)

func CreateTransferHandler(response http.ResponseWriter, request *http.Request) {

	MakeSureLoggedIn(response, request)

	data := struct {
		Errors []string
	}{
		Errors: []string{},
	}

	if request.Method == "POST" {
		CreateTransferPostHandler(response, request, &data)
	}

	render, err := template.ParseFiles("templates/create-transfer.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}

//Functie om de transfer op te pakken. Controleert of er voldoende balans is (canwithdrawfromaccount) en transfert vervolgens het geld dmv WithdrawFromAccount en AddToAccount
func CreateTransferPostHandler(response http.ResponseWriter, request *http.Request, data *struct{ Errors []string }) {
	request.ParseForm()

	id := helpers.GetAccountIdCookie(request)
	target_id := helpers.ConvertId(request)
	amount, _ := strconv.ParseFloat(request.FormValue("amount"), 64)
	description := (request.FormValue("description"))

	var canwithdraw bool = repositories.CanWithdrawFromAccount(id, amount)

	if canwithdraw {
		repositories.WithdrawFromAccount(id, amount, description)
		repositories.AddToAccount(target_id, amount, description)
		log.Println("Geld afgeschreven (van, naar, hoeveelheid, beschrijving): ", id, target_id, amount, description)

		http.Redirect(response, request, "/transfer-confirmed", http.StatusFound)
	} else {
		log.Println("Kon geen geld afschrijven")
		data.Errors = append(data.Errors, "Onvoldoende balans")
	}
}
