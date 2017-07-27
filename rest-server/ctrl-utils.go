package main

import (
	"os"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/oddbitsio/api/core"
)

type ICtrlUtils interface {
	Get(request *http.Request) map[string]string 
	WriteJsonResult(writer http.ResponseWriter, request *http.Request, result core.IModel, err error) 
}

type CtrlUtils struct { }

func (this *CtrlUtils) Get(request *http.Request) map[string]string {
	return mux.Vars(request);
}

func (this *CtrlUtils) WriteJsonResult(writer http.ResponseWriter, request *http.Request, result core.IModel, err error) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else if result.IsEmpty() {
		http.NotFound(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(result)
	}
}

type comboHandler func (writer http.ResponseWriter, request *http.Request)

func (this comboHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handlers.LoggingHandler(os.Stdout, http.HandlerFunc(this)).ServeHTTP(writer, request)
}