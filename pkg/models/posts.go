package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Title          string             `bson:"title,omitempty"`
	Description    string             `bson:"description,omitempty"`
	ImageUrl       string             `bson:"imageUrl,omitempty"`
	PlaylistUrl    string             `bson:"playlistUrl,omitempty"`
	Text           string             `bson:"text,omitempty"`
	FavouriteCount int                `bson:"favouriteCount,omitempty"`
	Creator        string             `bson:"creator,omitempty"`
	Mood           []string           `bson:"mood,omitempty"`
	Distance       int                `bson:"distance,omitempty"`
	Length         int                `bson:"length,omitempty"`
	Location       string             `bson:"location,omitempty"`
}

func GetAllPosts() *[]Post {
	var posts []Post
	coll := Client.Database("cp-server").Collection("posts")
	findOptions := options.Find()

	ctx := context.Background()

	cursor, err := coll.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var result Post
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal("cursor.Decode ERROR:", err)
		}
		posts = append(posts, result)
	}

	defer cursor.Close(ctx)

	return &posts
}
