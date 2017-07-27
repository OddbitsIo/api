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
}

func (this *OrganizationSvcMock) Get(code string) (*core.OrganizationModel, error) {
	return this.GetFunc(code)
}

func (this *OrganizationSvcMock) Save(organization *contracts.OrganizationModel) error {
	return this.SaveFunc(organization)
}

func TestGetOrganization_NotFound(t *testing.T) {
	var router = mux.NewRouter()
	ctrl := main.CreateOrganizationCtrl().RegisterRoutes(router)
	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(id string) (*contracts.OrganizationModel, error) {
			if (id != "123") {
				t.Errorf("Incorrect id: %s", id)
			}
            return &contracts.OrganizationModel {}, nil
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
			//Id: "2",
			Code: "Test",
			Name: "OrgName",
			TaxId: "OrgTaxId",
		}

	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(code string) (*core.OrganizationModel, error) {
			if code != "2" {
				t.Errorf("Inexpected code: %s", code)
			}
            return orgResult, nil
		}}
		
	request, _:= http.NewRequest("GET", "/organization/2", nil)
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