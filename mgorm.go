package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var errorWrapper = func(err error) error {
	return err
}

func SetErrorWrapper(fn func(err error) error) {
	errorWrapper = fn
}

type MgORM struct {
	db  *mongo.Database
	col *mongo.Collection
}

type Option struct {
	Collection string
	WrapError  func(err error) error
}

func NewORM(db *mongo.Database, collection string) *MgORM {
	return &MgORM{
		db:  db,
		col: db.Collection(collection),
	}
}

func (c *MgORM) Collection() *mongo.Collection {
	return c.col
}

func (c *MgORM) Find(ctx context.Context, filter interface{}) *FindBuilder {
	if filter == nil {
		filter = bson.M{}
	}
	return &FindBuilder{
		ctx:    ctx,
		col:    c.col,
		opt:    options.Find(),
		filter: filter,
	}
}

func (c *MgORM) Count(ctx context.Context, filter interface{}) (int64, error) {
	count, err := c.col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errorWrapper(err)
	}
	return count, nil
}

func (c *MgORM) Update(ctx context.Context, filter interface{}, update interface{}) *UpdateBuilder {
	return &UpdateBuilder{
		ctx:    ctx,
		opt:    options.Update(),
		col:    c.col,
		filter: filter,
		update: update,
	}
}

func (c *MgORM) Delete(ctx context.Context, filter interface{}) *DeleteBuilder {
	return &DeleteBuilder{
		ctx:    ctx,
		opt:    options.Delete(),
		col:    c.col,
		filter: filter,
	}
}

func (c *MgORM) Insert(ctx context.Context) *InsertBuilder {
	return &InsertBuilder{ctx: ctx, col: c.col}
}

func (c *MgORM) NewTransaction(ctx context.Context, callback func(tx mongo.Session) error) error {
	session, err := c.db.Client().StartSession()
	if err != nil {
		return errorWrapper(err)
	}

	err = callback(session)
	if err != nil {
		return errorWrapper(session.AbortTransaction(ctx))
	}
	return errorWrapper(session.CommitTransaction(ctx))
}

func (c *MgORM) CreateIndex(keys []string, name string, unique bool) error {
	opt := options.Index().SetUnique(unique)
	if name != "" {
		opt.SetName(name)
	}
	_, err := c.col.Indexes().CreateOne(Context(), mongo.IndexModel{
		Keys:    keys,
		Options: opt.SetBackground(true),
	})
	return errorWrapper(err)
}

func (c *MgORM) FindById(ctx context.Context, id string, result interface{}) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errorWrapper(err)
	}
	var filter = bson.M{"_id": objId}
	return errorWrapper(c.col.FindOne(ctx, filter).Decode(result))
}

func (c *MgORM) UpdateById(ctx context.Context, id string, update interface{}) (*mongo.UpdateResult, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errorWrapper(err)
	}
	var filter = bson.M{"_id": objId}
	result, err := c.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return result, nil
}

func (c *MgORM) DeleteById(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errorWrapper(err)
	}
	var filter = bson.M{"_id": objId}
	result, err := c.col.DeleteOne(ctx, filter)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return result, nil
}

func (c *MgORM) Aggregate(ctx context.Context) *AggregateBuilder {
	return &AggregateBuilder{
		ctx:      ctx,
		col:      c.col,
		pipeline: bson.A{},
	}
}
