package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateBuilder struct {
	ctx    context.Context
	multi  bool
	opt    *options.UpdateOptions
	col    *mongo.Collection
	filter interface{}
	update interface{}
	err    error
	result *mongo.UpdateResult
}

func (c *UpdateBuilder) Error() error {
	return c.err
}

func (c *UpdateBuilder) GetResult() *mongo.UpdateResult {
	return c.result
}

func (c *UpdateBuilder) SetUpsert(f bool) {
	c.opt.SetUpsert(f)
}

func (c *UpdateBuilder) Do() *UpdateBuilder {
	if c.multi {
		c.result, c.err = c.col.UpdateMany(c.ctx, c.filter, c.update, c.opt)
	} else {
		c.result, c.err = c.col.UpdateOne(c.ctx, c.filter, c.update, c.opt)
	}
	return c
}
