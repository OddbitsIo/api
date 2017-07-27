package core

type OrganizationModel struct {
	//Id string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"` 
	TaxId string `json:"taxId"`
}

func (this *OrganizationModel) IsEmpty() bool {
	return this.Code == ""
}