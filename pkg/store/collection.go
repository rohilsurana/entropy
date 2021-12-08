package store

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	collection *mongo.Collection
}

var (
	InsertFailedError = errors.New("failed to insert into db")
	UpdateFailedError = errors.New("failed to update in db")
	FindFailedError   = errors.New("failed to query db")
	NotFoundError     = errors.New("no such record(s) in db")
)

func (c *Collection) InsertOne(document interface{}) error {
	_, err := c.collection.InsertOne(context.TODO(), document)
	if err != nil {
		return fmt.Errorf("%w: %s", InsertFailedError, err)
	}
	return nil
}

func (c *Collection) UpdateByURN(urn string, document interface{}) error {
	_, err := c.collection.UpdateOne(context.TODO(),
		map[string]interface{}{"urn": urn},
		map[string]interface{}{"$set": document},
	)
	if err != nil {
		return fmt.Errorf("%w: %s", UpdateFailedError, err)
	}
	return nil
}

func (c *Collection) FindOne(filter interface{}, result interface{}) error {
	err := c.collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: %s", NotFoundError, err)
		}
		return fmt.Errorf("%w: %s", FindFailedError, err)
	}
	return nil
}

func (c *Collection) Find(filter map[string]interface{}, result interface{}) error {
	cursor, err := c.collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("%w: %s", FindFailedError, err)
	}
	err = cursor.All(context.TODO(), result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("%w: %s", NotFoundError, err)
		}
		return fmt.Errorf("%w: %s", FindFailedError, err)
	}
	return nil
}
