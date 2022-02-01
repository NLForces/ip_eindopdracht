package repositories

import (
	"main/types"
	. "main/types"
)

func GetAllAccounts() []Account {
	var accounts []Account

	connection().Find(&accounts)

	return accounts
}

func GetAccount(account_id uint) Account {
	var account Account

	connection().First(&account, account_id)

	return account
}

//Functie die controleert of het de ingevoerde pincode overeenkomt met die uit de database
func PincodeControle(account_id string, ingevoerdePincode string) bool {
	var databasePincode Account
	connection().Where("id = ?", account_id).Find(&databasePincode)

	if ingevoerdePincode == databasePincode.Pincode {
		return true
	}
	return false
}

//Maakt het account aan met de meegegeven gegevens
func CreateAccount(name string, code string, maxcredit float64, pincode string) {
	account := types.Account{Name: name, Code: code, Maxcredit: maxcredit, Pincode: pincode}
	connection().Create(&account)
}

//Controle of het account aangemaakt kan worden. Als de naam of code al bestaat, geeft hij een false terug
func CanCreateAccount(name string, code string) bool {

	alleAccounts := GetAllAccounts()

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
