package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/oddbitsio/api/contracts"
)

type IParamProvider interface {
	Get(request *http.Request) map[string]string	
}

type ParamProvider struct { }

func (this *ParamProvider) Get(request *http.Request) map[string]string {
	return mux.Vars(request);
}

type IResponseWriter interface {
	WriteJsonResult(writer http.ResponseWriter, request *http.Request, result contracts.IResult, err error)
}

type ResponseWriter struct { }

func (this *ResponseWriter) WriteJsonResult(writer http.ResponseWriter, request *http.Request, result contracts.IResult, err error) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	} else if result.IsEmpty() {
		http.NotFound(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(result)
	}
}