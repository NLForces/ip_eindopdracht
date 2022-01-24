package handlers

import (
	"net/http"
)

func LogOut(response http.ResponseWriter, account_id string) {
	http.SetCookie(response, nil)
}
