package repositories

import (
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
