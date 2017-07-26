package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/oddbitsio/api/contracts"
	"github.com/oddbitsio/api/services"
	"github.com/oddbitsio/api/mdb-repos"
)

type OrganizationCtrl struct {
	Service contracts.IOrganizationService
	Utils ICtrlUtils
}

func CreateOrganizationCtrl() *OrganizationCtrl {
	return &OrganizationCtrl {
		Utils: &CtrlUtils {},
		Service: &services.Organization {
			OrganizationRepo: &mdbrepos.Organization {}}}
}

func (this *OrganizationCtrl) RegisterRoutes(router *mux.Router) *OrganizationCtrl {
	router.Handle("/organization/{id}", comboHandler(this.GetOrganization)).Methods("GET")
	return this
}

func (this *OrganizationCtrl) GetOrganization(writer http.ResponseWriter, request *http.Request) {
	id := this.Utils.Get(request)["id"]
	org, err := this.Service.Get(id);
	this.Utils.WriteJsonResult(writer, request, org, err)
}