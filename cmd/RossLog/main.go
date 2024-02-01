package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rcampbell-sec/RossLogGo/internal/api"
	"github.com/rcampbell-sec/RossLogGo/internal/db"
	"github.com/rcampbell-sec/RossLogGo/internal/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// ae := api.GetAllEntries(ctx, entriesCollection)
	// for _, e := range ae {
	// 	fmt.Println(e.Title)
	// }

	ae := api.GetEntryByTitle("do", entriesCollection, ctx)
	for _, e := range ae {
		fmt.Println(e.Title)
	}

	myEntry := types.Entry{
		Title:     "test",
		Body:      "yeah",
		Tags:      []string{"1"},
		Datestamp: primitive.NewDateTimeFromTime(time.Now()),
	}
	result := api.InsertEntry(myEntry, entriesCollection, ctx)
	fmt.Printf("The insert returned %t\n", result)
}
