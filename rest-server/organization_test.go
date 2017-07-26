package main_test

import (
	"testing"
	"net/http"
	"errors"
	"strings"
	"net/http/httptest"
	"github.com/gorilla/mux"
	"github.com/oddbitsio/api/contracts"
	"github.com/oddbitsio/api/rest-server"
	"encoding/json"
)

type OrganizationSvcMock struct {
	GetFunc func(string) (*contracts.OrganizationResult, error)
}

func (this *OrganizationSvcMock) Get(id string) (*contracts.OrganizationResult, error) {
	return this.GetFunc(id)
}

func TestGetOrganization_NotFound(t *testing.T) {
	var router = mux.NewRouter()
	ctrl := main.CreateOrganizationCtrl().RegisterRoutes(router)
	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(id string) (*contracts.OrganizationResult, error) {
			if (id != "123") {
				t.Errorf("Incorrect id: %s", id)
			}
            return &contracts.OrganizationResult {}, nil
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
		GetFunc: func(id string) (*contracts.OrganizationResult, error) {
			if id != "456" {
				t.Errorf("Inexpected id: %s", id)
			}
            return &contracts.OrganizationResult {}, errors.New(errorMsg)
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
	orgResult := &contracts.OrganizationResult {
			Id: "2",
			Name: "OrgName",
			TaxId: "OrgTaxId"}

	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(id string) (*contracts.OrganizationResult, error) {
			if id != "2" {
				t.Errorf("Inexpected id: %s", id)
			}
            return orgResult, nil
		}}
		
	request, _:= http.NewRequest("GET", "/organization/2", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusOK {
        t.Errorf("Incorrect status code: %d", recorder.Code)
	}

	decodedResult := contracts.OrganizationResult { }
	json.NewDecoder(recorder.Body).Decode(&decodedResult)
	if decodedResult != *orgResult {
		t.Error("Bad message body")
	}
}