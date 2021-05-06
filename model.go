package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	col *mongo.Collection
}

func (c *Model) Collection() *mongo.Collection {
	return c.col
}

func (c *Model) Context() context.Context {
	return NewContext()
}

func (c *Model) Find(ctx context.Context, filter interface{}) *FindBuilder {
	if filter == nil {
		filter = bson.M{}
	}
	return &FindBuilder{
		ctx:    WithWrap(ctx),
		col:    c.col,
		opt:    options.Find(),
		filter: filter,
	}
}

func (c *Model) Count(ctx context.Context, filter interface{}) (int64, error) {
	count, err := c.col.CountDocuments(WithWrap(ctx), filter)
	if err != nil {
		return 0, errorWrapper(err)
	}
	return count, nil
}

func (c *Model) Exists(ctx context.Context, filter interface{}) bool {
	num, _ := c.col.CountDocuments(WithWrap(ctx), filter, options.Count().SetLimit(1))
	return num > 0
}

func (c *Model) Update(ctx context.Context, filter interface{}, update interface{}) *UpdateBuilder {
	return &UpdateBuilder{
		ctx:    WithWrap(ctx),
		opt:    options.Update(),
		col:    c.col,
		filter: filter,
		update: update,
	}
}

func (c *Model) Delete(ctx context.Context, filter interface{}) *DeleteBuilder {
	return &DeleteBuilder{
		ctx:    WithWrap(ctx),
		opt:    options.Delete(),
		col:    c.col,
		filter: filter,
	}
}

func (c *Model) Insert(ctx context.Context) *InsertBuilder {
	return &InsertBuilder{ctx: WithWrap(ctx), col: c.col}
}

func (c *Model) CreateIndex(keys []string, name string, unique bool) error {
	opt := options.Index().SetUnique(unique)
	if name != "" {
		opt.SetName(name)
	}

	var d = bson.D{}
	for _, key := range keys {
		d = append(d, bson.E{Key: key, Value: 1})
	}
	_, err := c.col.Indexes().CreateOne(c.Context(), mongo.IndexModel{
		Keys:    d,
		Options: opt.SetBackground(true),
	})
	return errorWrapper(err)
}

func (c *Model) FindById(ctx context.Context, id string, result interface{}) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errorWrapper(err)
	}
	var filter = bson.M{"_id": objId}
	return errorWrapper(c.col.FindOne(WithWrap(ctx), filter).Decode(result))
}

func (c *Model) UpdateById(ctx context.Context, id string, update interface{}) (*mongo.UpdateResult, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errorWrapper(err)
	}
	var filter = bson.M{"_id": objId}
	result, err := c.col.UpdateOne(WithWrap(ctx), filter, update)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return result, nil
}

func (c *Model) DeleteById(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errorWrapper(err)
	}
	var filter = bson.M{"_id": objId}
	result, err := c.col.DeleteOne(WithWrap(ctx), filter)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return result, nil
}

func (c *Model) Aggregate(ctx context.Context) *AggregateBuilder {
	return &AggregateBuilder{
		ctx:      WithWrap(ctx),
		col:      c.col,
		pipeline: bson.A{},
	}
}
