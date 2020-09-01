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

func (c *MgORM) FindAll(filter interface{}) *FindBuilder {
	if filter == nil {
		filter = bson.M{}
	}
	opt := options.Find()
	return &FindBuilder{
		col:    c.col,
		opt:    opt,
		filter: filter,
	}
}

func (c *MgORM) Count(filter interface{}) (int64, error) {
	return c.col.CountDocuments(newContext(), filter)
}
