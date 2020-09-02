package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgORM struct {
	db  *mongo.Database
	col *mongo.Collection
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
	return c.col.CountDocuments(ctx, filter)
}

func (c *MgORM) UpdateMany(ctx context.Context, filter interface{}, update interface{}) *UpdateBuilder {
	return &UpdateBuilder{
		ctx:    ctx,
		multi:  true,
		opt:    options.Update(),
		col:    c.col,
		filter: filter,
		update: update,
	}
}

func (c *MgORM) UpdateOne(ctx context.Context, filter interface{}, update interface{}) *UpdateBuilder {
	return &UpdateBuilder{
		ctx:    ctx,
		multi:  false,
		opt:    options.Update(),
		col:    c.col,
		filter: filter,
		update: update,
	}
}

func (c *MgORM) DeleteMany(ctx context.Context, filter interface{}) *DeleteBuilder {
	return &DeleteBuilder{
		ctx:    ctx,
		multi:  true,
		opt:    options.Delete(),
		col:    c.col,
		filter: filter,
	}
}

func (c *MgORM) DeleteOne(ctx context.Context, filter interface{}) *DeleteBuilder {
	return &DeleteBuilder{
		ctx:    ctx,
		multi:  false,
		opt:    options.Delete(),
		col:    c.col,
		filter: filter,
	}
}

func (c *MgORM) InsertOne(ctx context.Context, document interface{}) (primitive.ObjectID, error) {
	result, err := c.col.InsertOne(ctx, document)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (c *MgORM) InsertMany(ctx context.Context, documents []interface{}) ([]primitive.ObjectID, error) {
	results, err := c.col.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}

	var ids = make([]primitive.ObjectID, 0)
	for _, item := range results.InsertedIDs {
		ids = append(ids, item.(primitive.ObjectID))
	}
	return ids, nil
}

func (c *MgORM) NewTransaction(ctx context.Context, callback func(tx mongo.Session) error) error {
	session, err := c.db.Client().StartSession()
	if err != nil {
		return err
	}

	err = callback(session)
	if err != nil {
		return session.AbortTransaction(ctx)
	}
	return session.CommitTransaction(ctx)
}

func (c *MgORM) CreateIndex(keys []string, name string, unique bool) error {
	opt := options.Index().SetUnique(unique)
	if name != "" {
		opt.SetName(name)
	}
	_, err := c.col.Indexes().CreateOne(Context(), mongo.IndexModel{
		Keys:    keys,
		Options: opt,
	})
	return err
}
