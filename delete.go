package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeleteBuilder struct {
	ctx    context.Context
	filter interface{}
	opt    *options.DeleteOptions
	col    *mongo.Collection
}

func (c *DeleteBuilder) One() (*mongo.DeleteResult, error) {
	return c.col.DeleteOne(c.ctx, c.filter, c.opt)
}

func (c *DeleteBuilder) All() (*mongo.DeleteResult, error) {
	return c.col.DeleteMany(c.ctx, c.filter, c.opt)
}
