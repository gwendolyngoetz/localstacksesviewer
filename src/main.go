package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gwendolyngoetz/localstacksesviewer/handlers"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5001", "")
	flag.Parse()
	http.HandleFunc("/checkhealth", handlers.HandlerCheckHealth)
	http.HandleFunc("/emails", handlers.HandlerGetEmails)
	http.HandleFunc("/email", handlers.HandlerGetEmail)

	fmt.Printf("Listening on %s\n", *listenAddr)
	http.ListenAndServe(*listenAddr, nil)
}
