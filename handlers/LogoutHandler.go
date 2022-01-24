package handlers

import (
	"log"
	"net/http"
)

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("Logged out")

	http.SetCookie(response, nil)
	http.Redirect(response, request, "/login", http.StatusFound)
}
