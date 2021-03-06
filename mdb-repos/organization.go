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

type OrganizationRepo struct { }

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
	session := db.getSession();
	defer session.Close()

	collection := session.DB("oddbits").C("organizations")

	result := OrganizationDto{}
	if err := collection.Find(bson.M{"code": code}).One(&result); err != nil {
		return nil, err
	}
	return this.toModel(&result), nil
}

func (this *OrganizationRepo) Save(model *core.OrganizationModel) error  {
	session := db.getSession();
	session.SetMode(mgo.Strong, false)
	defer session.Close()

	collection := session.DB("oddbits").C("organizations")

	index := mgo.Index {
		Key:        []string{ "code" },
		Unique:     true,
	}
	
	if err := collection.EnsureIndex(index); err != nil {
		return err
	}

	dto := this.toDto(model)

	var selector map[string]interface{}
	if dto.Id.Valid() {
		selector = bson.M{"_id": dto.Id}
	} else {
		selector = bson.M{"code": dto.Code}
	}
	_, err := collection.Upsert(selector, &dto)

	return err
}

func (this *OrganizationRepo) Delete(code string) error  {
	session := db.getSession();
	session.SetMode(mgo.Eventual, false)
	defer session.Close()

	collection := session.DB("oddbits").C("organizations")

	return collection.Remove(bson.M{"code": code})
}

