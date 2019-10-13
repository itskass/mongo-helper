package mongohelper

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type Doc interface {
	GetID() string
}

// Save will insert or replace a document in a given collection.
// If a document with the same _id exists the objet will be replaced.
// Otherwise the document will be inserted
func Save(ctx context.Context, col *mongo.Collection, doc Doc) error {
	id, err := primitive.ObjectIDFromHex(doc.GetID())
	if err != nil {
		return err
	}

	// attempt find and replace
	if err := Replace(ctx, col, doc, bson.M{"_id": id}); err == nil {
		return nil
	}

	// attempt to insert
	_, err = col.InsertOne(ctx, doc)
	return err
}
