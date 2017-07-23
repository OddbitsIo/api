package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

)

type ICtrlUtils interface {
	GetParams(request *http.Request) map[string] string
	HandleGetResult(writer http.ResponseWriter, request *http.Request, result interface{}, err error) 
}

type CtrlUtils struct {

}

func (this *CtrlUtils) GetParams(request *http.Request) map[string] string {
	return mux.Vars(request)
}

func (this *CtrlUtils) HandleGetResult(writer http.ResponseWriter, request *http.Request, result interface{}, err error) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	} else if result == nil {
		http.NotFound(writer, request)
	} else {
		json.NewEncoder(writer).Encode(result)
	}
}

