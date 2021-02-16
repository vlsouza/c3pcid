package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Oie Nara!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	if len(process.env.PORT) > 0 {
		port := process.env.PORT
	} else {
		port := ":8080"
	}
	log.Fatal(http.ListenAndServe(port, router))
}
