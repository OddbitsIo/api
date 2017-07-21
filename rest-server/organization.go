package main

import (
	"net/http"
	"github.com/oddbitsio/api/contracts"
	"github.com/oddbitsio/api/services"
	"github.com/oddbitsio/api/mdb-repos"
)

type OrganizationCtrl struct {
	OrganizationService contracts.IOrganizationService
}

func CreateOrganizationCtrl() *OrganizationCtrl {
	return &OrganizationCtrl { 
		OrganizationService : &services.Organization {
			OrganizationRepo : &mdbrepos.Organization {}}}
}

func (ctrl *OrganizationCtrl) GetOrganization(writer http.ResponseWriter, request *http.Request) {
	params := GetParams(request)
	id := params["id"]
	org, err := ctrl.OrganizationService.Get(id);
	HandleGetResult(writer, request, org, err)
}