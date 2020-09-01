package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgORM struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewORM(db *mongo.Database) *MgORM {
	return &MgORM{
		db: db,
	}
}

func (c *MgORM) Collection(name string) *MgORM {
	c.col = c.db.Collection(name)
	return c
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
