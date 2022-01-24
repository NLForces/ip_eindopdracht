package repositories

import (
	"main/helpers"
	. "main/types"
	"math"
)

// Script om alle transacties op te vragen per account ID. Saldo wordt opgeteld of afgetrokken
// Menno heeft dit uitgelegd: https://imgur.com/qyAvmD2
func GetSaldoForAccount(account_id uint) float64 {
	var totaalTransactions []Transaction
	connection().Where("account_id = ?", account_id).Find(&totaalTransactions)

	var saldo float64

	for _, value := range totaalTransactions {
		if value.Type == "debet" {
			saldo += value.Amount
		} else {
			saldo -= value.Amount
		}
	}

	saldo = math.Round(saldo*100) / 100 //Math.round functie om af te ronden op 2 cijfers achter de komma

	return saldo
}

//Onderstaande script lijkt op GetSaldoForAccount. Aangevuld met Gorm: https://gorm.io/docs/query.html)
func GetTransactionsForAccount(account_id uint) []Transaction {
	var totaalTransactions []Transaction
	connection().Where("account_id = ?", account_id).Find(&totaalTransactions)

	return totaalTransactions
}

//Controleert of het ingevulde amount van het saldo afgehaald kan worden.
//Maakt ook gebuik van persoondata.MaxCredit (opdracht 7)
func CanWithdrawFromAccount(account_id uint, amount float64) bool {

	var persoonData Account
	connection().Where("id = ?", account_id).Find(&persoonData)

	saldo := GetSaldoForAccount(account_id)
	canWithdraw := persoonData.Maxcredit <= saldo-amount

	return canWithdraw
}

//Pusht de geüpdatede records terug naar de database wanneer er een geld af moet
func WithdrawFromAccount(account_id uint, amount float64, description string) Transaction {

	user := Transaction{AccountId: account_id, Amount: amount, Type: "credit", Description: description}
	connection().Create(&user)

	return user
}

//Pusht de records naar de database wanneer er geld bij moet
func AddToAccount(account_id uint, amount float64, description string) Transaction {

	user := Transaction{AccountId: account_id, Amount: amount, Type: "debet", Description: description}
	connection().Create(&user)

	return user
}

//Ontvangt de filepath en schrijft (als er een error is) de output van een JSON file weg naar de database
//de vars in de struct blijven in deze func en trekken hun gegevens uit de genoemde plek in de JSON.
//Bron: Opdracht 5 van college 5
func ImportTransactions(filepath string) error {

	type alleTransacties struct {
		AccountId   uint    `json:"id"`
		Amount      float64 `json:"total_amount"`
		Type        string  `json:"transaction_type"`
		Description string  `json:"message"`
	}

	var jsondata []alleTransacties

	jsonContinue := helpers.LoadData(filepath, &jsondata)
	if jsonContinue != nil {
		return jsonContinue
	}
	var enkeleTransactie Transaction

	for _, v := range jsondata {
		enkeleTransactie = Transaction{AccountId: v.AccountId, Amount: v.Amount, Type: v.Type, Description: v.Description}
		connection().Create(&enkeleTransactie)
	}

	return nil
}
