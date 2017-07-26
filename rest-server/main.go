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
	router.Handle("/organization/{id}", comboHandler(orgCtrl.GetOrganization)).Methods("GET")
	
	log.Printf("Listening on :%d", httpPort)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router))
}

type comboHandler func (writer http.ResponseWriter, request *http.Request)

func (this comboHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handlers.LoggingHandler(os.Stdout, http.HandlerFunc(this)).ServeHTTP(writer, request)
}
