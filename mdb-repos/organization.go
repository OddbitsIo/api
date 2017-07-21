package mdbrepos

import (
	"github.com/oddbitsio/api/core"
)

type Organization struct {

}

func (repo *Organization) Get(id string) (*core.OrganizationDto, error) {
	dummy := core.OrganizationDto{ Id: id }
	return &dummy, nil
}

func (repo *Organization) Save(organization *core.OrganizationDto) error {
	return nil
}