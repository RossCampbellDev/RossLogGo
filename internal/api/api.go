package api

import (
	"context"
	"fmt"

	"github.com/rcampbell-sec/RossLogGo/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllEntries(coll *mongo.Collection, ctx context.Context) []types.Entry {
	var results []types.Entry

	allEntries, err := coll.Find(ctx, bson.D{}) // context, filter, options (not included here)

	if err != nil {
		panic(err)
	}
	defer allEntries.Close(ctx)

	for allEntries.Next(ctx) {
		var e types.Entry

		if err := allEntries.Decode(&e); err != nil {
			panic(err)
		}
		results = append(results, e)
	}

	return results
}

func GetEntryByTitle(title string, coll *mongo.Collection, ctx context.Context) []types.Entry {
	title = fmt.Sprintf(".*%s.*", title)

	var results []types.Entry
	var filter = bson.M{"title": bson.M{"$regex": title}}

	allEntries, err := coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	defer allEntries.Close(ctx)

	for allEntries.Next(ctx) {
		var e types.Entry

		if err := allEntries.Decode(&e); err != nil {
			panic(err)
		}
		results = append(results, e)
	}

	return results
}

func InsertEntry(e types.Entry, coll *mongo.Collection, ctx context.Context) bool {
	result, err := coll.InsertOne(ctx, e)
	if err != nil {
		return false
	}

	if _, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return false
	}
	return true
}
