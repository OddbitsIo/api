package main_test

import (
	"bytes"
	"testing"
	"net/http"
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
		Utils: &main.CtrlUtils{}}
}

func TestGetOrganization_NotFound(t *testing.T) {

	ctrl := createOrganizationTestCtrl()
	ctrl.Service = &OrganizationSvcMock {
		GetFunc: func(id string) (*contracts.OrganizationResult, error) {
            return nil, nil
		}}
		
	request, _:= http.NewRequest("GET", "/test",
		bytes.NewBuffer([]byte(`{"id":"99"}`)))
    recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ctrl.GetOrganization)
	
    handler.ServeHTTP(recorder, request)

    if recorder.Code != http.StatusNotFound {
        t.Errorf("Incorrect status code: %d", recorder.Code)
    }
}