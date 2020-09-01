package mgorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgORM struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewMgORM(db *mongo.Database) *MgORM {
	return &MgORM{
		db: db,
	}
}

func (c *MgORM) Collection(name string) *MgORM {
	c.col = c.db.Collection(name)
	return c
}

func (c *MgORM) Find(filter interface{}) *FindBuilder {
	if filter == nil {
		filter = bson.M{}
	}
	return &FindBuilder{
		col:    c.col,
		opt:    options.Find(),
		filter: filter,
	}
}

func (c *MgORM) Count(filter interface{}) (int64, error) {
	return c.col.CountDocuments(newContext(), filter)
}

func (c *MgORM) UpdateMany(filter interface{}, update interface{}) *UpdateBuilder {
	return &UpdateBuilder{
		multi:  true,
		opt:    options.Update(),
		col:    c.col,
		filter: filter,
		update: update,
	}
}

func (c *MgORM) UpdateOne(filter interface{}, update interface{}) *UpdateBuilder {
	return &UpdateBuilder{
		multi:  false,
		opt:    options.Update(),
		col:    c.col,
		filter: filter,
		update: update,
	}
}

func (c *MgORM) DeleteMany(filter interface{}) *DeleteBuilder {
	return &DeleteBuilder{
		multi:  true,
		opt:    options.Delete(),
		col:    c.col,
		filter: filter,
	}
}

func (c *MgORM) DeleteOne(filter interface{}) *DeleteBuilder {
	return &DeleteBuilder{
		multi:  false,
		opt:    options.Delete(),
		col:    c.col,
		filter: filter,
	}
}
