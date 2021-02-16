package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Oie Nara!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(port, router))
}
