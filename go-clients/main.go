package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

var (
	port     = 27018
	host     = "localhost"
	database = "tests"
)

func insertUser(db *mongo.Database) {
	user := bson.M{
		"name":   "yuta",
		"prefer": bson.A{"golang", "kubernetes", "statistics", "SRE"},
		"role": bson.D{
			{Key: "enable", Value: true},
			{Key: "name", Value: "admin"},
		},
	}
	res, err := db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(res)
}

func insertUsers(db *mongo.Database) {
	users := []interface{}{
		bson.M{
			"name":   "aaa",
			"prefer": bson.A{"ruby", "design"},
			"role": bson.D{
				{Key: "enable", Value: false},
				{Key: "name", Value: "dev"},
			},
		},
		bson.M{
			"name":   "bbb",
			"prefer": bson.A{"python", "sre"},
			"role": bson.D{
				{Key: "enable", Value: true},
				{Key: "name", Value: "dev"},
			},
		},
	}
	res, err := db.Collection("users").InsertMany(context.Background(), users)
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(res)
}

func deleteUsers(db *mongo.Database) {
	coll := db.Collection("users")
	err := coll.Drop(context.Background())
	if err != nil {
		panic(err.Error())
	}
}

func connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dsn := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongo.Connect(ctx, dsn)
	if err != nil {
		panic(err.Error())
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err.Error())
	}
	return client
}

func main() {
	client := connect()
	db := client.Database(database)

	cmd := "connect"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "connect":
		fmt.Println("connect OK")
	case "insertUser":
		insertUser(db)
	case "insertUsers":
		insertUsers(db)
	case "deleteUsers":
		deleteUsers(db)
	}
}
