package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/oddbitsio/api/core"
	"github.com/oddbitsio/api/services"
	"github.com/oddbitsio/api/mdb-repos"
)

type IOrganizationSvc interface {
	Get(code string) (*core.OrganizationModel, error)
	Save(organization *core.OrganizationModel) error
	Delete(code string) error
}

type OrganizationCtrl struct {
	Service IOrganizationSvc
	Utils ICtrlUtils
}

func CreateOrganizationCtrl() *OrganizationCtrl {
	return &OrganizationCtrl {
		Utils: &CtrlUtils { },
		Service: &services.OrganizationSvc {
			OrganizationRepo: &mdbrepos.OrganizationRepo { },
		},
	}
}

func (this *OrganizationCtrl) RegisterRoutes(router *mux.Router) *OrganizationCtrl {
	router.Handle("/organization/{code}", comboHandler(this.Get)).Methods("GET")
	router.Handle("/organization", comboHandler(this.Save)).Methods("POST")
	router.Handle("/organization/{code}", comboHandler(this.Delete)).Methods("DELETE")
	return this
}

func (this *OrganizationCtrl) Get(writer http.ResponseWriter, request *http.Request) {
	code := this.Utils.GetParams(request)["code"]
	org, err := this.Service.Get(code);
	this.Utils.WriteJsonResult(writer, request, org, err)
}

func (this *OrganizationCtrl) Save(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
    var orgPayload core.OrganizationModel
    if err := decoder.Decode(&orgPayload); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if err := this.Service.Save(&orgPayload); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	org, err := this.Service.Get(orgPayload.Code);
	this.Utils.WriteJsonResult(writer, request, org, err)
}

func (this *OrganizationCtrl) Delete(writer http.ResponseWriter, request *http.Request) {
	code := this.Utils.GetParams(request)["code"]
	if err := this.Service.Delete(code); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}