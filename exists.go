package mdbh

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Exists(ctx context.Context, col *mongo.Collection,
	query bson.M) bool {
	count, err := col.CountDocuments(ctx, query)
	if err != nil {
		return false
	}
	return count > 0
}
