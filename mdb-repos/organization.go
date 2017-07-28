package mdbrepos

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/oddbitsio/api/core"
)


type OrganizationDto struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Code string `bson:"code"`
	Name string `bson:"name"`
	TaxId string `bson:"taxId"`
}

type OrganizationRepo struct {

}

func (this *OrganizationRepo) toModel(dto *OrganizationDto) *core.OrganizationModel {
	return &core.OrganizationModel {
		Id: dto.Id.Hex(),
		Code: dto.Code,
		Name: dto.Name,
		TaxId: dto.TaxId,
	}
}

func (this *OrganizationRepo) toDto(model *core.OrganizationModel) *OrganizationDto {
	dto := &OrganizationDto {
		Code: model.Code,
		Name: model.Name,
		TaxId: model.TaxId,
	}

	if model.Id != "" {
		dto.Id = bson.ObjectIdHex(model.Id)
	}

	return dto
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
		Key:        []string{ "code" },
		Unique:     true,
	}
	
	if err = collection.EnsureIndex(index); err != nil {
		return err
	}

	dto := this.toDto(model)

	var selector map[string]interface{}
	if dto.Id.Valid() {
		selector = bson.M{"_id": dto.Id}
	} else {
		selector = bson.M{"code": dto.Code}
	}
	_, err = collection.Upsert(selector, &dto)

	return err
}