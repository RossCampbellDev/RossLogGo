package main

import (
	"context"
	"fmt"

	"github.com/rcampbell-sec/RossLogGo/db"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ctx = context.TODO()
)

func main() {
	client := db.GetMongoClient(ctx)
	defer client.Disconnect(ctx)

	entries := db.GetEntryCollection(client)

	var results []bson.M
	allEntries, err := entries.Find(ctx, bson.D{}) // context, filter, options (not included here)
	if err != nil {
		panic(err)
	}

	for allEntries.Next(ctx) {
		var item bson.M
		if err := allEntries.Decode(&item); err != nil {
			panic(err)
		}
		results = append(results, item)
	}
	allEntries.Close(ctx)

	fmt.Println(results)
}
