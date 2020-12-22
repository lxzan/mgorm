package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AggregateBuilder struct {
	ctx      context.Context
	col      *mongo.Collection
	pipeline bson.A
}

func (c *AggregateBuilder) Match(match interface{}) *AggregateBuilder {
	var p = bson.M{
		"$match": match,
	}
	c.pipeline = append(c.pipeline, p)
	return c
}

func (c *AggregateBuilder) Project(project interface{}) *AggregateBuilder {
	var p = bson.M{
		"$project": project,
	}
	c.pipeline = append(c.pipeline, p)
	return c
}

func (c *AggregateBuilder) Limit(limit int64) *AggregateBuilder {
	var p = bson.M{
		"$limit": limit,
	}
	c.pipeline = append(c.pipeline, p)
	return c
}

func (c *AggregateBuilder) Skip(skip int64) *AggregateBuilder {
	var p = bson.M{
		"$skip": skip,
	}
	c.pipeline = append(c.pipeline, p)
	return c
}

func (c *AggregateBuilder) Group(key string, m bson.M) *AggregateBuilder {
	m["_id"] = "$" + key
	var m1 = bson.M{
		"$group": m,
	}
	c.pipeline = append(c.pipeline, m1)
	return c
}

func (c *AggregateBuilder) Sort(sort interface{}) *AggregateBuilder {
	var p = bson.M{
		"$sort": sort,
	}
	c.pipeline = append(c.pipeline, p)
	return c
}

func (c *AggregateBuilder) All(results interface{}) error {
	cursor, err := c.col.Aggregate(c.ctx, c.pipeline)
	if err != nil {
		return errorWrapper(err)
	}
	return errorWrapper(cursor.All(c.ctx, results))
}
