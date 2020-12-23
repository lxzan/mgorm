package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InsertBuilder struct {
	ctx context.Context
	col *mongo.Collection
}

func (c *InsertBuilder) One(document interface{}) (interface{}, error) {
	result, err := c.col.InsertOne(c.ctx, document)
	if err != nil {
		return primitive.ObjectID{}, errorWrapper(err)
	}
	return result.InsertedID, nil
}

func (c *InsertBuilder) All(documents []interface{}) ([]interface{}, error) {
	results, err := c.col.InsertMany(c.ctx, documents)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return results.InsertedIDs, nil
}
