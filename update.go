package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateBuilder struct {
	ctx    context.Context
	opt    *options.UpdateOptions
	col    *mongo.Collection
	filter interface{}
	update interface{}
}

func (c *UpdateBuilder) SetUpsert(f bool) *UpdateBuilder {
	c.opt.SetUpsert(f)
	return c
}

func (c *UpdateBuilder) One() (*mongo.UpdateResult, error) {
	result, err := c.col.UpdateOne(c.ctx, c.filter, c.update, c.opt)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return result, nil
}

func (c *UpdateBuilder) All() (*mongo.UpdateResult, error) {
	result, err := c.col.UpdateMany(c.ctx, c.filter, c.update, c.opt)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return result, nil
}

func (c *UpdateBuilder) SetOne() (*mongo.UpdateResult, error) {
	result, err := c.col.UpdateOne(c.ctx, c.filter, bson.M{"$set": c.update}, c.opt)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return result, nil
}

func (c *UpdateBuilder) SetAll() (*mongo.UpdateResult, error) {
	result, err := c.col.UpdateMany(c.ctx, c.filter, bson.M{"$set": c.update}, c.opt)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return result, nil
}
