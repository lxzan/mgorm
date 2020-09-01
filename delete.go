package mgorm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeleteBuilder struct {
	multi  bool
	filter interface{}
	opt    *options.DeleteOptions
	col    *mongo.Collection
	err    error
	result *mongo.DeleteResult
}

func (c *DeleteBuilder) Error() error {
	return c.err
}

func (c *DeleteBuilder) GetResult() *mongo.DeleteResult {
	return c.result
}

func (c *DeleteBuilder) Do() *DeleteBuilder {
	if c.multi {
		c.result, c.err = c.col.DeleteMany(newContext(), c.filter, c.opt)
	} else {
		c.result, c.err = c.col.DeleteOne(newContext(), c.filter, c.opt)
	}
	return c
}
