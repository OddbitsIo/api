package main_test

import (
	"testing"
	"net/http"
	"errors"
	"strings"
	"net/http/httptest"
	"github.com/gorilla/mux"
	"github.com/oddbitsio/api/core"
	"github.com/oddbitsio/api/rest-server"
	"encoding/json"
)

type OrganizationSvcMock struct {
	GetFunc func(string) (*core.OrganizationModel, error)
	SaveFunc func(organization *core.OrganizationModel) error
	DeleteFunc func(string) error
}

func (this *OrganizationSvcMock) Get(code string) (*core.OrganizationModel, error) {
	return this.GetFunc(code)
}

func (this *OrganizationSvcMock) Save(organization *core.OrganizationModel) error {
	return this.SaveFunc(organization)
}

func (this *OrganizationSvcMock) Delete(code string) error {
	return this.Delete(code)
}

func TestGetOrganization_NotFound(t *testing.T) {
	var router = mux.NewRouter()
	ctrl := main.CreateOrganizationCtrl().RegisterRoutes(router)
	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(id string) (*core.OrganizationModel, error) {
			if (id != "123") {
				t.Errorf("Incorrect code: %s", id)
			}
            return &core.OrganizationModel {}, nil
		}}
		
	request, _:= http.NewRequest("GET", "/organization/123", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusNotFound {
        t.Errorf("Incorrect status code: %d", recorder.Code)
    }
}

func TestGetOrganization_Error(t *testing.T) {
	errorMsg := "suprise!"

	var router = mux.NewRouter()
	ctrl := main.CreateOrganizationCtrl().RegisterRoutes(router)
	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(code string) (*core.OrganizationModel, error) {
			if code != "456" {
				t.Errorf("Inexpected code: %s", code)
			}
            return &core.OrganizationModel {}, errors.New(errorMsg)
		}}
		
	request, _:= http.NewRequest("GET", "/organization/456", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusBadRequest {
        t.Errorf("Incorrect status code: %d", recorder.Code)
	}
	var body = strings.TrimSpace(recorder.Body.String())
	if body != errorMsg {
		t.Errorf("Invalid response body: %s", body)
	}
}

func TestGetOrganization_Success(t *testing.T) {
	var router = mux.NewRouter()
	ctrl := main.CreateOrganizationCtrl().RegisterRoutes(router)
	orgResult := &core.OrganizationModel {
		Code: "Test",
		Name: "OrgName",
		TaxId: "OrgTaxId",
	}

	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(code string) (*core.OrganizationModel, error) {
			if code != "Test" {
				t.Errorf("Inexpected code: %s", code)
			}
            return orgResult, nil
		}}
		
	request, _:= http.NewRequest("GET", "/organization/Test", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusOK {
        t.Errorf("Incorrect status code: %d", recorder.Code)
	}

	decodedResult := core.OrganizationModel { }
	json.NewDecoder(recorder.Body).Decode(&decodedResult)
	if decodedResult != *orgResult {
		t.Error("Bad message body")
	}
}