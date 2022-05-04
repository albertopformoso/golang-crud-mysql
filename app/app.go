package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	//Check if the .env file exists and load it
	LoadDotEnv(".env")
	port := ":" + os.Getenv("PORT")

	// Initialize multiplexor
	router := mux.NewRouter().StrictSlash(true)

	// Handle Routes
	router.HandleFunc("/", index)
	router.HandleFunc("/create", createEmployee)

	// Initialize server
	fmt.Printf("Serving on: http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

