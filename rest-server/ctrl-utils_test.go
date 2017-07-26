package main_test

import (
	"net/http"
)

type ParamProviderMock struct {
	GetResult map[string]string
}

func (this *ParamProviderMock) Get(request *http.Request) map[string]string {
	return this.GetResult
}
