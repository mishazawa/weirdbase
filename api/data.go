package api

import (
	"os"
	"context"
	"time"
	"fmt"
	"log"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Connection mongo.Client


func CreateRecord (data map[string]interface{}) error {
	collection := Connection.Database("testing").Collection("test")

	// TODO return BSON
	bin := InsertRandomData(data)

	_, err := collection.InsertOne(context.TODO(), bin)

	if err != nil {
		return err
	}

	return nil
}

func InsertRandomData (data map[string]interface {}) []byte {
	// for k, v := range data {
		// TODO вскрыть себе вены
		// reflect kind of value
	// }
	return make([]byte, 0)
}

func Connect () error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	Connection, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)))

	if err != nil {
		return err
	}

	defer disconnect(ctx)

	log.Println("Connected to", os.Getenv("DB_HOST"))

	ctx, cancel = context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	return Connection.Ping(ctx, readpref.Primary())
}

func disconnect (ctx context.Context) {
	if err := Connection.Disconnect(ctx); err != nil {
		panic(err)
	}
}
