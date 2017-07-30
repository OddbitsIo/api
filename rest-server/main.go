package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/oddbitsio/api/services"
)

func main() {
	httpPort := 9000
	router := mux.NewRouter()
	
	if err := services.Init(); err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialize dependencies: %s", err.Error()))
	}
	defer services.TearDown()

	CreateOrganizationCtrl().RegisterRoutes(router)

	log.Printf("Listening on :%d", httpPort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router))
}