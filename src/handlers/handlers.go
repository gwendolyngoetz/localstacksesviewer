package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gwendolyngoetz/localstacksesviewer/types"
)

func getEmails() types.SesResponse {
	res, err := http.Get("http://localhost:4566/_aws/ses")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var response types.SesResponse

	err = decoder.Decode(&response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

func getEmailById(id string) types.SesResponse {
	res, err := http.Get("http://localhost:4566/_aws/ses?id=" + id)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var response types.SesResponse

	err = decoder.Decode(&response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

func HandlerCheckHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi\n")
}

func HandlerGetEmails(w http.ResponseWriter, r *http.Request) {
	response := getEmails()
	tmpl := template.Must(template.ParseFiles("templates/main.html"))
	tmpl.Execute(w, response)
}

func HandlerGetEmail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	response := getEmailById(id)
	tmpl := template.Must(template.ParseFiles("templates/email.html"))
	tmpl.Execute(w, response.Messages[0])
}
