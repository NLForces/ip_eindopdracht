package handlers

import (
	"fmt"
	"html/template"
	"log"
	"main/repositories"
	"net/http"
	"strconv"
)

func CreationConfirmedHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("CreationConfirmed")

	allAccounts := repositories.GetAllAccounts()
	account_id := strconv.FormatUint(uint64(allAccounts[len(allAccounts)-1].ID), 10) //Roept het ID op van het laatste account en convert dit in een string
	name := allAccounts[len(allAccounts)-1].Name
	code := allAccounts[len(allAccounts)-1].Code
	maxcredit := allAccounts[len(allAccounts)-1].Maxcredit

	var info string = fmt.Sprintf("Account Id: "+account_id+"\n"+"Naam: "+name+"\n"+"Code: "+code+"\n"+"Rood: "+"%f", account_id, maxcredit)

	// data := struct {
	// 	tekst []string
	// }{
	// 	tekst: info,
	// }

	render, err := template.ParseFiles("templates/account-created.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, info)
}
