package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/oddbitsio/api/core"
	"github.com/oddbitsio/api/services"
)

type IOrganizationSvc interface {
	Get(code string) (*core.OrganizationModel, error)
	Save(organization *core.OrganizationModel) error
	Delete(code string) error
}

type OrganizationCtrl struct {
	Service IOrganizationSvc
	utils ctrlUtils
}

func CreateOrganizationCtrl() *OrganizationCtrl {
	return &OrganizationCtrl {
		utils: ctrlUtils { },
		Service: services.CreateOrganizaitonSvc(),
	}
}

func (this *OrganizationCtrl) RegisterRoutes(router *mux.Router) *OrganizationCtrl {
	router.Handle("/organization/{code}", comboHandler(this.get)).Methods("GET")
	router.Handle("/organization", comboHandler(this.save)).Methods("POST")
	router.Handle("/organization/{code}", comboHandler(this.delete)).Methods("DELETE")
	return this
}

func (this *OrganizationCtrl) get(writer http.ResponseWriter, request *http.Request) {
	code := this.utils.getParams(request)["code"]
	org, err := this.Service.Get(code);
	this.utils.writeJsonResult(writer, request, org, err)
}

func (this *OrganizationCtrl) save(writer http.ResponseWriter, request *http.Request) {
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
	this.utils.writeJsonResult(writer, request, org, err)
}

func (this *OrganizationCtrl) delete(writer http.ResponseWriter, request *http.Request) {
	code := this.utils.getParams(request)["code"]
	if err := this.Service.Delete(code); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}