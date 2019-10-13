package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(ctx context.Context, col *mongo.Collection, doc Doc, query bson.M) error {
	res := col.FindOne(ctx, query)
	if res.Err() != nil {
		return res.Err()
	}
	return res.Decode(doc)
}
