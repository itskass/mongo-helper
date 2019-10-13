package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Replace(ctx context.Context, col *mongo.Collection, doc Doc,
	query bson.M) error {

	res, err := col.ReplaceOne(ctx, query, doc)
	if err != nil {
		return err
	}
	if res.ModifiedCount < 1 {
		return errors.New("nothing modified")
	}
	return nil
}
