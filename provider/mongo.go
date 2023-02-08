package provider

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Mongo() (client *mongo.Client, err error) {
	host := os.Getenv("MONGO_HOST")
	user := os.Getenv("MONGO_USER")
	pwd := os.Getenv("MONGO_PWD")

	credential := options.Credential{
		Username: user,
		Password: pwd,
	}
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", host)).SetAuth(credential)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		_ = fmt.Errorf("%s", err)
	}

	return
}
