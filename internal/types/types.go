package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entry struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title" form:"title"`
	Body      string             `bson:"body" form:"body"`
	Tags      []string           `bson:"tags" form:"tags"`
	Datestamp primitive.DateTime `bson:"datestamp"`
}

type User struct {
	ID       string `bson:"id,omitempty"`
	Username string `bson:"username"`
	Passhash string `bson:"passhash"`
}
