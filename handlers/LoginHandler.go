package handlers

import (
	"html/template"
	"log"
	"main/helpers"
	"main/repositories"
	"net/http"
)

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
		data.Errors = append(data.Errors, "Wrong pincode")
	}
}
