package routing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func RunServer() {
	client := db.GetMongoClient(ctx)
	defer client.Disconnect(ctx)
	entriesCollection = db.GetEntryCollection(client)

	router := gin.Default()
	router.GET("/logs", getAllEntries)
	router.GET("/logs/:title", getEntryByTitle)
	router.POST("/new", insertEntry)

	router.Run("localhost:8080")
	fmt.Println("--- Server Running ---")
}

func getAllEntries(c *gin.Context) {
	entries := api.GetAllEntries(ctx, entriesCollection)
	if len(entries) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no entries found"})
		return
	}
	c.IndentedJSON(http.StatusOK, entries)
}

func getEntryByTitle(c *gin.Context) {
	title := c.Param("title")
	result := api.GetEntryByTitle(ctx, entriesCollection, title)
	if len(result) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no matches for " + title})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func insertEntry(c *gin.Context) {
	var myEntry types.Entry
	if err := c.Bind(&myEntry); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not bind entry"})
		return
	}

	myEntry.Datestamp = primitive.NewDateTimeFromTime(time.Now())

	result := api.InsertEntry(ctx, entriesCollection, myEntry)

	if !result {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to insert entry"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
