package resource

import (
	"context"
	"fmt"

	"github.com/odpf/entropy/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


const RepositoryName = "resource"

type Repository struct {
	DB *mongo.Database
}

func (rc *Repository) Create(r *model.Resource) (*model.Resource, error) {
	coll := rc.DB.Collection(RepositoryName)
	findRes, err := rc.GetResourceByURN(r.URN)
	if err == nil && findRes != nil {
		return nil, fmt.Errorf("resource with URN %s already exists.", r.URN)
	}
	res, err := coll.InsertOne(context.TODO(), r)
	if err != nil {
		return nil, err
	}
	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		r.ID = id.Hex()
	}
	return r, nil
}

func (rc *Repository) Update(r *model.Resource) (*model.Resource, error) {
	coll := rc.DB.Collection(RepositoryName)
	_, err := coll.UpdateByID(context.TODO(), r.ID, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (rc *Repository) GetResourceByURN(urn string) (*model.Resource, error) {
	coll := rc.DB.Collection(RepositoryName)
	res := model.Resource{}
	err := coll.FindOne(context.TODO(), bson.D{{Key: "urn", Value: urn}}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (rc *Repository) GetResources(kind string, parent string) ([]*model.Resource, error) {
	var res []*model.Resource
	coll := rc.DB.Collection(RepositoryName)
	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "kind", Value: kind}, {Key: "parent", Value: parent}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var r model.Resource
		if err := cursor.Decode(&r); err != nil {
			return nil, err
		}
		res = append(res, &r)
	}
	return res, nil
}

func (rc *Repository) Delete(urn string) error { return nil }
