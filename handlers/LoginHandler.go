package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	"net/http"
)

// DISCLAIMER: The login mechanism is NOT representative for a real-world scenario.
// It is only used to provide a login-like functionality without the complexity that
// usualy comes with features like this.

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("LoginHandler")

	data := struct {
		Errors []string
	}{
		Errors: []string{},
	}

	if request.Method == "POST" {
		handlePost(response, request, &data)
	}

	render, err := template.ParseFiles("templates/login.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}

//Functie toegevoegd van pincodeklopt, haalt de gegevens uit het formulier die hij doorstuurt.
func handlePost(response http.ResponseWriter, request *http.Request, data *struct{ Errors []string }) {
	request.ParseForm()

	var account_id = request.FormValue("account_id")
	var pincodeKlopt = repositories.PincodeControle(account_id, request.FormValue("pincode"))

	helpers.SetAccountIdCookie(response, account_id)

	if pincodeKlopt {
		http.Redirect(response, request, "/welcome", http.StatusFound)
	} else {
		data.Errors = append(data.Errors, "Onjuiste pincode")
	}
}

// func CreateWithdrawalPostHandler(response http.ResponseWriter, request *http.Request, data *struct{ Errors []string }) {
// 	request.ParseForm()

// 	id := helpers.GetAccountIdCookie(request)
// 	amount, _ := strconv.ParseFloat(request.FormValue("amount"), 64)
// 	description := (request.FormValue("description"))

// 	var canwithdraw bool = repositories.CanWithdrawFromAccount(id, amount)

// 	if canwithdraw {
// 		repositories.CanWithdrawFromAccount(id, amount)
// 		repositories.WithdrawFromAccount(id, amount, description)
// 		http.Redirect(response, request, "/withdrawal-confirmed", http.StatusFound)
// 	} else {
// 		data.Errors = append(data.Errors, "Onvoldoende balans")
// 	}
// }
