package handlers

import (
	"html/template"
	"io"
	"log"
	"main/repositories"
	"net/http"
	"os"
)

func ImportTransactionsHandler(response http.ResponseWriter, request *http.Request) {
	log.Println("ImportTransactionsHandler")

	MakeSureLoggedIn(response, request)

	data := struct {
		Errors []string
	}{
		Errors: []string{},
	}

	if request.Method == "POST" {
		ImportTransactionsPostHandler(response, request, &data)
	}

	render, err := template.ParseFiles("templates/import-transactions.html")

	if err != nil {
		log.Println(err)
	}

	render.Execute(response, data)
}

func ImportTransactionsPostHandler(response http.ResponseWriter, request *http.Request, data *struct{ Errors []string }) {
	request.ParseMultipartForm(10 << 20)

	// Process the file

	file, handler, err := request.FormFile("file")

	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)

		data.Errors = append(data.Errors, "The uploaded file could not be processed")

		return
	}

	// Upload the file

	dst, err := os.Create("./data/" + handler.Filename)
	defer dst.Close()

	if err != nil {
		data.Errors = append(data.Errors, "The uploaded file could not be processed")
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		data.Errors = append(data.Errors, "The uploaded file could not be processed")
		return
	}

	//Handler die importtransactions aanroept en het filepath meegeeft van de JSON file. Stuurt je na afloop door naar de /transactions
	jsonContinue := repositories.ImportTransactions(dst.Name())
	if jsonContinue != nil {
		data.Errors = append(data.Errors, "Error loading JSON")
	}

	log.Println("Json file imported")

	http.Redirect(response, request, "/transactions", http.StatusFound)
}
