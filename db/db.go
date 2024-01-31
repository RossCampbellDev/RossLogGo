package db

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ENTRY_COLLECTION = "EntryCollection"
	USER_COLLECTION  = "UserCollection"

	MONGO_USER        string
	MONGO_PASS        string
	MONGO_DB_NAME     string
	MONGO_CONN_STRING string
)

// dont use short declaration!!!! think of the scope
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	MONGO_USER = os.Getenv("MONGO_USER")
	MONGO_PASS = os.Getenv("MONGO_PASS")
	MONGO_DB_NAME = os.Getenv("MONGO_DB_NAME")
	MONGO_CONN_STRING = os.Getenv("MONGO_CONN_STRING")

	MONGO_CONN_STRING = strings.Replace(MONGO_CONN_STRING, "<username>", MONGO_USER, 1)
	MONGO_CONN_STRING = strings.Replace(MONGO_CONN_STRING, "<password>", MONGO_PASS, 1)
}

func GetMongoClient(ctx context.Context) *mongo.Client {
	fmt.Printf("conn: %s\n\n", MONGO_CONN_STRING)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_CONN_STRING).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	// test db connection
	// if err := client.Database(MONGO_DB_NAME).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	// 	panic(err)
	// }

	fmt.Println("Connected to Mongo Client")
	return client
}

func GetEntryCollection(c *mongo.Client) *mongo.Collection {
	return c.Database(MONGO_DB_NAME).Collection(ENTRY_COLLECTION)
}

func GetUserCollection(c *mongo.Client) *mongo.Collection {
	return c.Database(MONGO_DB_NAME).Collection(USER_COLLECTION)
}
