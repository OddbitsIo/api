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
	
	orgCtrl := createOrganizationCtrl()
	router.Handle("/organization/{id}", getHandler(orgCtrl.GetOrganization)).Methods("GET")
	

	log.Printf("Listening on :%d", httpPort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router))
}

func getHandler(routeHandler func(writer http.ResponseWriter, request *http.Request)) http.Handler {
	return handlers.LoggingHandler(os.Stdout, http.HandlerFunc(routeHandler))
}
