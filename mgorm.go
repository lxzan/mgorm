package mgorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var errorWrapper = func(err error) error {
	return err
}

func SetErrorWrapper(fn func(err error) error) {
	errorWrapper = fn
}

type MgORM struct {
	db *mongo.Database
}

func NewORM(db *mongo.Database) *MgORM {
	return &MgORM{db: db}
}

func (c *MgORM) NewModel(collection string) *Model {
	return &Model{
		col: c.db.Collection(collection),
	}
}

func (c *MgORM) NewTransaction(ctx context.Context, callback func(tx mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	ctx = WithWrap(ctx)
	session, err := c.db.Client().StartSession()
	if err != nil {
		return nil, errorWrapper(err)
	}
	return session.WithTransaction(ctx, callback)
}
