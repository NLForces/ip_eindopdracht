package main

import (
	"log"
	"net/http"

	"main/handlers"
	"main/repositories"

	"github.com/joho/godotenv"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://ip-eindopdracht-bram.nl/login", 301)
}

func main() {
	log.Println("Starting application")

	// Dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database
	if repositories.Connected() {
		log.Println("Starting database connection")
	}
	http.HandleFunc("/", redirect)

	http.HandleFunc("/welcome", handlers.ShowWelcomeHandler)

	http.HandleFunc("/login", handlers.LoginHandler)

	http.HandleFunc("/saldo", handlers.ShowSaldoHandler)
	http.HandleFunc("/transactions", handlers.ListTransactionsHandler)

	http.HandleFunc("/withdrawal", handlers.CreateWithdrawalHandler)
	http.HandleFunc("/withdrawal-confirmed", handlers.WithdrawalConfirmedHandler)

	http.HandleFunc("/transfer", handlers.CreateTransferHandler)
	http.HandleFunc("/transfer-confirmed", handlers.TransferConfirmedHandler)

	http.HandleFunc("/import-transactions", handlers.ImportTransactionsHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
