package main

import (
	"context"
	"github.com/lxzan/mgorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CatchError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:admin@192.168.1.203:27017/casbin?authSource=admin"))
	CatchError(err)
	CatchError(client.Connect(context.Background()))

	var results = make([]struct {
		Id    string `json:"id" bson:"_id"`
		Total int64  `json:"total" bson:"total"`
	}, 0)
	orm := mgorm.NewORM(client.Database("casbin"), "casbin_rule")
	err = orm.
		Aggregate(context.Background()).
		Group("v2", bson.M{
			"total": bson.M{"$sum": 1},
		}).
		Sort(bson.M{"_id": -1}).
		All(&results)
	println(&err)
}
