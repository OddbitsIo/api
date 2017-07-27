package mdbrepos

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/oddbitsio/api/core"
)


type OrganizationDto struct {
	//Id string
	Code string `bson:"code"`
	Name string `bson:"name"`
	TaxId string `bson:"taxId"`
}

type OrganizationRepo struct {

}

func (this *OrganizationRepo) toModel(dto *OrganizationDto) *core.OrganizationModel {
	return &core.OrganizationModel {
		Code: dto.Code,
		Name: dto.Name,
		TaxId: dto.TaxId,
	}
}

func (this *OrganizationRepo) toDto(model *core.OrganizationModel) *OrganizationDto {
	return &OrganizationDto {
		Code: model.Code,
		Name: model.Name,
		TaxId: model.TaxId,
	}
}

func (this *OrganizationRepo) Get(code string) (*core.OrganizationModel, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return nil, err
	}

	defer session.Close()
	
	credentials := mgo.Credential {
		Username: "oddbits",
		Password: "yellowcamelridesbike",
	}
	if err = session.Login(&credentials); err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("oddbits").C("organizations")

	result := OrganizationDto{}
	if err = collection.Find(bson.M{"code": code}).One(&result); err != nil {
		return nil, err
	}
	return this.toModel(&result), nil
}

func (this *OrganizationRepo) Save(model *core.OrganizationModel) error  {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return err
	}

	defer session.Close()

	credentials := mgo.Credential {
		Username: "oddbits",
		Password: "yellowcamelridesbike",
	}
	if err = session.Login(&credentials); err != nil {
		return err
	}

	session.SetMode(mgo.Strong, true)

	collection := session.DB("oddbits").C("organizations")

	index := mgo.Index {
		Key:        []string{ "Code" },
		Unique:     true,
	}
	
	if err = collection.EnsureIndex(index); err != nil {
		return err
	}

	dto := this.toDto(model)
	_, err = collection.Upsert(bson.M{"Code": dto.Code}, &dto)

	return err
}