package handlers

import (
	"html/template"
	"log"
	"main/repositories"
	"net/http"
	"strconv"
)

//Uitbreiding idee van Menno
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

func CreateAccountPostHandler(response http.ResponseWriter, request *http.Request, data *struct{ Errors []string }) {
	request.ParseForm()

	name := (request.FormValue("name"))
	code := (request.FormValue("code"))
	maxcredit, _ := strconv.ParseFloat(request.FormValue("maxcredit"), 64)
	pincode := (request.FormValue("pincode"))

	var cancreate bool = CanCreateAccount(name, code)

	if cancreate {
		repositories.CreateAccount(name, code, maxcredit, pincode)
		CreationConfirmed(name, code, maxcredit, response)
		http.Redirect(response, request, "/account-created.html", http.StatusFound)
		log.Println("Account aangemaakt")
	} else {
		log.Println("Kon geen account aanmaken")
		data.Errors = append(data.Errors, "Kon geen account aanmaken")
	}
}

func CanCreateAccount(name string, code string) bool {

	alleAccounts := repositories.GetAllAccounts()

	for _, value := range alleAccounts {
		if value.Name == name {
			return false
		}
		if value.Code == code {
			return false
		}

	}

	return true
}

func CreationConfirmed(name string, code string, maxcredit float64, response http.ResponseWriter) {
	log.Println("CreationConfirmed")

	allAccounts := repositories.GetAllAccounts()
	account_id := allAccounts[len(allAccounts)-1].ID //Roept het ID op van het laatste account

	data := struct {
		name       string
		account_id uint
		code       string
		maxcredit  float64
	}{
		name:       name,
		account_id: account_id,
		code:       code,
		maxcredit:  maxcredit,
	}

	render, err := template.ParseFiles("templates/account-created.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}
