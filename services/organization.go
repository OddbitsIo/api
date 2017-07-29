package services

import (
	"github.com/oddbitsio/api/core"
)

type IOrganizationRepo interface {
	Get(code string) (*core.OrganizationModel, error)
	Save(organization *core.OrganizationModel) error
	Delete(code string) error
}

type OrganizationSvc struct {
	OrganizationRepo IOrganizationRepo
}

func (this *OrganizationSvc) Get(code string) (*core.OrganizationModel, error) {
	org, err := this.OrganizationRepo.Get(code)
	return org, err
}

func (this *OrganizationSvc) Save(organization *core.OrganizationModel) error  {
	return this.OrganizationRepo.Save(organization)
}

func (this *OrganizationSvc) Delete(code string) error {
	return this.OrganizationRepo.Delete(code)
}