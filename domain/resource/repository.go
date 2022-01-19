package resource

import (
	"errors"
	"fmt"

	"github.com/odpf/entropy/domain/model"
	"github.com/odpf/entropy/pkg/store"
)

const RepositoryName = "resource"

type Repository struct {
	DB *store.DB
}

var (
	ResourceAlreadyExistsError = errors.New("resource already exists")
	NoResourceFoundError       = errors.New("no resource(s) found")
)

func (rc *Repository) Create(r *model.Resource) (*model.Resource, error) {
	coll := rc.DB.GetCollection(RepositoryName)
	err := coll.FindOne(map[string]interface{}{"urn": r.URN}, &model.Resource{})
	if err != nil {
		if !errors.Is(err, store.NotFoundError) {
			return nil, fmt.Errorf("%w: URN = %s", ResourceAlreadyExistsError, r.URN)
		}
		return nil, err
	}
	err = coll.InsertOne(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (rc *Repository) Update(r *model.Resource) (*model.Resource, error) {
	coll := rc.DB.GetCollection(RepositoryName)
	err := coll.UpdateOne(map[string]interface{}{"urn": r.URN}, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (rc *Repository) GetResourceByURN(urn string) (*model.Resource, error) {
	coll := rc.DB.GetCollection(RepositoryName)
	res := model.Resource{}
	err := coll.FindOne(map[string]interface{}{"urn": urn}, &res)
	if err != nil {
		if errors.Is(err, store.NotFoundError) {
			return nil, fmt.Errorf("%w: URN = %s", NoResourceFoundError, urn)
		}
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
