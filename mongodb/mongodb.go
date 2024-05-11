package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProvideMongoDBCli(uri string) (*mongo.Client, error) {
	return mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(uri),
	)
}
