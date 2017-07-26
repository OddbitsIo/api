package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	
)

func main() {
	httpPort := 9000
	router := mux.NewRouter()
	
	CreateOrganizationCtrl().RegisterRoutes(router)

	log.Printf("Listening on :%d", httpPort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router))
}