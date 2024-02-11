package api

import (
	"context"
	"fmt"

	"github.com/rcampbell-sec/RossLogGo/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllEntries(ctx context.Context, coll *mongo.Collection) ([]types.Entry, error) {
	// var results []types.Entry
	/*
		create a slice of length zero, rather than a "nil slice
		either way works, but this is more explicit.  with the above line, the empty slice
		is initialised when we try to append anyway
	*/
	results := make([]types.Entry, 0)

	allEntries, err := coll.Find(ctx, bson.D{}) // context, filter, options (not included here)

	if err != nil {
		return nil, fmt.Errorf("Error when searching DB for records: %w", err)
	}
	defer allEntries.Close(ctx)

	for allEntries.Next(ctx) {
		var e types.Entry

		if err := allEntries.Decode(&e); err != nil {
			return nil, fmt.Errorf("Failed to decode records into Entry format: %w", err)
		}
		results = append(results, e)
	}

	return results, nil
}

func GetEntriesByTitle(ctx context.Context, coll *mongo.Collection, title string) ([]types.Entry, error) {
	title = fmt.Sprintf(".*%s.*", title)

	results := make([]types.Entry, 0)
	var filter = bson.M{"title": bson.M{"$regex": title}}

	allEntries, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("Error when searching DB for records: %w", err)
	}
	defer allEntries.Close(ctx)

	for allEntries.Next(ctx) {
		var e types.Entry

		if err := allEntries.Decode(&e); err != nil {
			return nil, fmt.Errorf("Failed to decode records into Entry format: %w", err)
		}
		results = append(results, e)
	}

	return results, nil
}

func InsertEntry(ctx context.Context, coll *mongo.Collection, e types.Entry) bool {
	result, err := coll.InsertOne(ctx, e)
	if err != nil {
		return false
	}

	if _, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return false
	}
	return true
}
