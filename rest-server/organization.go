package main

import (
	"net/http"
	"github.com/oddbitsio/api/contracts"
	"github.com/oddbitsio/api/services"
	"github.com/oddbitsio/api/mdb-repos"
)

type OrganizationCtrl struct {
	Service contracts.IOrganizationService
	ParamProvider IParamProvider
	ResponseWriter IResponseWriter
}

func createOrganizationCtrl() *OrganizationCtrl {
	return &OrganizationCtrl {
		ParamProvider: &ParamProvider {},
		ResponseWriter: &ResponseWriter {},
		Service: &services.Organization {
			OrganizationRepo: &mdbrepos.Organization {}}}
}

func (this *OrganizationCtrl) GetOrganization(writer http.ResponseWriter, request *http.Request) {
	id := this.ParamProvider.Get(request)["id"]
	org, err := this.Service.Get(id);
	this.ResponseWriter.WriteJsonResult(writer, request, org, err)
}