package main_test

import (
	"testing"
	"net/http"
	"errors"
	"strings"
	"net/http/httptest"
	"github.com/oddbitsio/api/contracts"
	"github.com/oddbitsio/api/rest-server"
)

type OrganizationSvcMock struct {
	GetFunc func(string) (*contracts.OrganizationResult, error)
}

func (this *OrganizationSvcMock) Get(id string) (*contracts.OrganizationResult, error) {
	return this.GetFunc(id)
}

func createOrganizationTestCtrl() *main.OrganizationCtrl {
	return &main.OrganizationCtrl { 
		ResponseWriter: &main.ResponseWriter { }}
}

func TestGetOrganization_NotFound(t *testing.T) {
	orgId := "123"

	ctrl := createOrganizationTestCtrl()
	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(id string) (*contracts.OrganizationResult, error) {
			if id != orgId {
				t.Errorf("Inexpected id: %s", id)
			}
            return &contracts.OrganizationResult {}, nil
		}}
	ctrl.ParamProvider = &ParamProviderMock {
		GetResult: map[string] string { "id": orgId } }
		
	request, _:= http.NewRequest("GET", "/test", nil)
	recorder := httptest.NewRecorder()
	ctrl.GetOrganization(recorder, request)

    if recorder.Code != http.StatusNotFound {
        t.Errorf("Incorrect status code: %d", recorder.Code)
    }
}

func TestGetOrganization_Error(t *testing.T) {
	orgId := "123"
	errorMsg := "suprise!"
	ctrl := createOrganizationTestCtrl()
	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(id string) (*contracts.OrganizationResult, error) {
			if id != orgId {
				t.Errorf("Inexpected id: %s", id)
			}
            return &contracts.OrganizationResult {}, errors.New(errorMsg)
		}}
	ctrl.ParamProvider = &ParamProviderMock {
		GetResult: map[string] string { "id": orgId } }
		
	request, _:= http.NewRequest("GET", "/test", nil)
	recorder := httptest.NewRecorder()
	ctrl.GetOrganization(recorder, request)

    if recorder.Code != http.StatusBadRequest {
        t.Errorf("Incorrect status code: %d", recorder.Code)
	}
	var body = strings.TrimSpace(recorder.Body.String())
	if body != errorMsg {
		t.Errorf("Invalid response body: %s", body)
	}
}