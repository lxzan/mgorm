package main

import (
	"context"
	"github.com/lxzan/mgorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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

	orm := mgorm.NewORM(client.Database("casbin"), "operate_log")

	var results = make([]struct {
		Id primitive.ObjectID `json:"id" bson:"_id"`
	}, 0)

	q := orm.Find(mgorm.WithTimeout(time.Second), nil).Limit(3).All(&results)
	println(&q)
}
