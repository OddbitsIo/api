package contracts

type OrganizationResult struct {
	Id string `json:"id"`
	Name string `json:"name"` 
	TaxId string `json:"taxId"`
}

func (this *OrganizationResult) IsEmpty() bool {
	return this.Id == ""
}

type IOrganizationService interface {
	Get(id string) (*OrganizationResult, error)
}