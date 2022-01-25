package handlers

import (
	"html/template"
	"log"
	"main/repositories"
	"net/http"
	"strconv"
)

//Uitbreiding idee van Menno. In feite CreateTransferHandler gekopiëerd en aangepast
//Roept createposthandler aan als er iets is ingevuld in de 'create account'
func CreateAccountHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("CreateWithdrawalHandler")

	data := struct {
		Errors []string
	}{
		Errors: []string{},
	}

	if request.Method == "POST" {
		CreateAccountPostHandler(response, request, &data)
	}

	render, err := template.ParseFiles("templates/create-account.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}

//
func CreateAccountPostHandler(response http.ResponseWriter, request *http.Request, data *struct{ Errors []string }) {
	request.ParseForm()

	name := (request.FormValue("name"))
	code := "IBAN " + (request.FormValue("code"))
	maxcredit, _ := strconv.ParseFloat(request.FormValue("maxcredit"), 64)
	pincode := (request.FormValue("pincode"))

	var cancreate bool = repositories.CanCreateAccount(name, code)

	if cancreate {
		repositories.CreateAccount(name, code, maxcredit, pincode)
		http.Redirect(response, request, "/creation-confirmed", http.StatusFound)
		log.Println("Account aangemaakt")
	} else {
		log.Println("Kon geen account aanmaken")
		data.Errors = append(data.Errors, "Kon geen account aanmaken")
	}
}
