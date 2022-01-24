package helpers

import (
	"net/http"
	"strconv"
	"time"
)

func SetAccountIdCookie(response http.ResponseWriter, account_id string) {
	cookie := http.Cookie{
		Name:    "account_id",
		Value:   account_id,
		Expires: time.Now().Add(10 * time.Minute), //Zet de expiration date van de cookie op 3 minuten
	}

	http.SetCookie(response, &cookie)
}

func LogOut(response http.ResponseWriter) {
	http.SetCookie(response, nil)
}

func GetAccountIdCookie(request *http.Request) uint {
	var _cookie, err = request.Cookie("account_id")

	if err != nil {
		return 0
	}

	var value, _ = strconv.ParseUint(_cookie.Value, 10, 32)

	return uint(value)
}

//Functie om ingevoerd ID te converteren naar een UINT
func ConvertId(request *http.Request) uint {
	form := request.FormValue("target_account_id")

	var value, _ = strconv.ParseUint(form, 10, 32)

	return uint(value)
}
