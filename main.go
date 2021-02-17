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

	if os.Getenv("PORT") != "" { 
		port = os.Getenv(“PORT”) 
	  } else { 
		port = "8080" 
	  }
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port),nil, router))
}
