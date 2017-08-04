package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/oddbitsio/api/services"
)

func main() {
	
	port := os.Getenv("API_PORT")
	if (port == "") {
		panic("API_PORT is not defined")
	}

	router := mux.NewRouter()
	
	if err := services.Init(); err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialize dependencies: %s", err.Error()))
	}
	defer services.TearDown()

	CreateOrganizationCtrl().RegisterRoutes(router)

	log.Printf("Listening on :%s", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}