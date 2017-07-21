package contracts

type OrganizationResult struct {
	Id string `json:"id"`
	Name string `json:"name"` 
	TaxId string `json:"taxId"`
}

type IOrganizationService interface {
	Get(id string) (*OrganizationResult, error)
}