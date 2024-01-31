package main

import (
	"context"
	"fmt"

	"github.com/rcampbell-sec/RossLogGo/db"
	"github.com/rcampbell-sec/RossLogGo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx               = context.TODO()
	entriesCollection *mongo.Collection
)

func main() {
	client := db.GetMongoClient(ctx)
	defer client.Disconnect(ctx)
	entriesCollection = db.GetEntryCollection(client)

	// ae := getAllEntries()
	// for _, e := range ae {
	// 	fmt.Println(e.Title)
	// }

	ae := getEntryByTitle("do")
	for _, e := range ae {
		fmt.Println(e.Title)
	}
}

func getAllEntries() []types.Entry {
	var results []types.Entry

	allEntries, err := entriesCollection.Find(ctx, bson.D{}) // context, filter, options (not included here)

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

func getEntryByTitle(title string) []types.Entry {
	title = fmt.Sprintf(".*%s.*", title)

	var results []types.Entry
	var filter = bson.M{"title": bson.M{"$regex": title}}

	allEntries, err := entriesCollection.Find(ctx, filter)
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
