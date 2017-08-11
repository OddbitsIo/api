package main

import (
	"os"
	"fmt"
	"strings"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
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
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), corsHandler(router)))
}

func corsHandler(router *mux.Router) http.Handler {
	var envCorsOrigins = os.Getenv("API_ALLOWED_CORS_ORIGINS");
	var allowedOrigins []string
	if (envCorsOrigins == "") {
		allowedOrigins = []string { "*" }
	} else {
		allowedOrigins = strings.Split(envCorsOrigins, ",")
	}

	return handlers.CORS(handlers.AllowedOrigins(allowedOrigins))(router)
}