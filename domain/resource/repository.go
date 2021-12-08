package resource

import (
	"fmt"

	"github.com/odpf/entropy/domain/model"
	"github.com/odpf/entropy/pkg/store"
	"go.mongodb.org/mongo-driver/bson"
)

const RepositoryName = "resource"

type Repository struct {
	DB *store.DB
}

func (rc *Repository) Create(r *model.Resource) (*model.Resource, error) {
	coll := rc.DB.GetCollection(RepositoryName)
	findRes, err := rc.GetResourceByURN(r.URN)
	if err == nil && findRes != nil {
		return nil, fmt.Errorf("resource with URN %s already exists.", r.URN)
	}
	err = coll.InsertOne(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (rc *Repository) Update(r *model.Resource) (*model.Resource, error) {
	coll := rc.DB.GetCollection(RepositoryName)
	err := coll.UpdateByURN(r.URN, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (rc *Repository) GetResourceByURN(urn string) (*model.Resource, error) {
	coll := rc.DB.GetCollection(RepositoryName)
	res := model.Resource{}
	err := coll.FindOne(bson.D{{Key: "urn", Value: urn}}, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (rc *Repository) GetResources(kind string, parent string) ([]*model.Resource, error) {
	var res []*model.Resource
	coll := rc.DB.GetCollection(RepositoryName)
	err := coll.Find(map[string]interface{}{"kind": kind, "parent": parent}, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rc *Repository) Delete(urn string) error { return nil }
