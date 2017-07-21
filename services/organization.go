package services

import (
	"github.com/oddbitsio/api/core"
	"github.com/oddbitsio/api/contracts"
)


type Organization struct {
	OrganizationRepo core.IOrganizationRepository
}

func (service *Organization) Get(id string) (*contracts.OrganizationResult, error) {
	org, err := service.OrganizationRepo.Get(id)
	if (err != nil) {
		return nil, err
	}

	result := &contracts.OrganizationResult {
		Id: org.Id,
		Name: org.Name,
		TaxId: org.TaxId }
		
	return result, nil
}