package main

import (
	"os"
	"fmt"
	"log"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	httpPort := 9000
	router := mux.NewRouter()
	
	orgCtrl := CreateOrganizationCtrl()
	router.Handle("/organization/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(orgCtrl.GetOrganization))).Methods("GET")
	


	log.Printf("Listening on :%d", httpPort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router))
}
