package main

import (
	"net/http"
	"github.com/oddbitsio/api/contracts"
	"github.com/oddbitsio/api/services"
	"github.com/oddbitsio/api/mdb-repos"
)

type OrganizationCtrl struct {
	Service contracts.IOrganizationService
	Utils ICtrlUtils
}

func createOrganizationCtrl() *OrganizationCtrl {
	return &OrganizationCtrl {
		Utils: &CtrlUtils{},
		Service: &services.Organization {
			OrganizationRepo: &mdbrepos.Organization {}}}
}

func (this *OrganizationCtrl) GetOrganization(writer http.ResponseWriter, request *http.Request) {
	params := this.Utils.GetParams(request)
	id := params["id"]
	org, err := this.Service.Get(id);
	this.Utils.HandleGetResult(writer, request, org, err)
}