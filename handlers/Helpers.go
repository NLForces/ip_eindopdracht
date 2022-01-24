package handlers

import (
	"fmt"
	"main/helpers"
	"net/http"
)

func MakeSureLoggedIn(response http.ResponseWriter, request *http.Request) {
	id := helpers.GetAccountIdCookie(request)

	if id == 0 {
		http.Redirect(response, request, "/login", http.StatusFound)
	} else {
		helpers.SetAccountIdCookie(response, fmt.Sprint(id)) //Als er activiteit is en de cookie is niet verlopen, ververst hij de cookie
	}
}
