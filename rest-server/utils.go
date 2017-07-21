package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetParams(request *http.Request) map[string] string {
	return mux.Vars(request)
}

func HandleGetResult(writer http.ResponseWriter, request *http.Request, result interface{}, err error) {
	if err != nil {
		http.Error(writer, err.Error(), 500)
	} else if result == nil {
		http.NotFound(writer, request)
	} else {
		json.NewEncoder(writer).Encode(result)
	}
}
