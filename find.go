package mgorm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindBuilder struct {
	col    *mongo.Collection
	opt    *options.FindOptions
	filter interface{}
	err    error
}

func (c *FindBuilder) Error() error {
	return c.err
}

func (c *FindBuilder) Select(fields interface{}) *FindBuilder {
	c.opt.SetProjection(fields)
	return c
}

func (c *FindBuilder) Offset(num int64) *FindBuilder {
	c.opt.SetSkip(num)
	return c
}

func (c *FindBuilder) Limit(num int64) *FindBuilder {
	c.opt.SetLimit(num)
	return c
}

func (c *FindBuilder) Sort(sort interface{}) *FindBuilder {
	c.opt.SetSort(sort)
	return c
}

func (c *FindBuilder) All(v interface{}) *FindBuilder {
	cursor, err := c.col.Find(newContext(), c.filter, c.opt)
	if err != nil {
		c.err = err
		return c
	}
	c.err = cursor.All(newContext(), v)
	return c
}
