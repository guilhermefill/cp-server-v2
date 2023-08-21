package models

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Title          string             `bson:"title" json:"title"`
	Description    string             `bson:"description" json:"description"`
	ImageUrl       string             `bson:"imageUrl" json:"imageUrl"`
	PlaylistUrl    string             `bson:"playlistUrl" json:"playlist_url"`
	Text           string             `bson:"text" json:"text"`
	FavouriteCount int                `bson:"favouriteCount" json:"favourite_count"`
	Creator        string             `bson:"creator" json:"creator"`
	Mood           []string           `bson:"mood" json:"mood"`
	Distance       int                `bson:"distance" json:"distance"`
	Length         int                `bson:"length" json:"length"`
	Location       string             `bson:"location" json:"location"`
}

func GetAllPosts() *[]Post {
	var posts []Post
	coll := Client.Database("cp-server").Collection("posts")
	findOptions := options.Find()

	ctx := context.Background()

	cursor, err := coll.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.New(os.Stdout, "coll.Find ERROR:"+err.Error()+"\t"+"", 1)
		return nil
	}

	for cursor.Next(ctx) {
		var result Post
		err := cursor.Decode(&result)
		if err != nil {
			log.New(os.Stdout, "cursor.Decode ERROR:"+err.Error()+"\t"+"", 1)
			return nil
		}
		posts = append(posts, result)
	}

	defer cursor.Close(ctx)

	return &posts
}

func GetPostById(id string) *Post {
	var post Post
	coll := Client.Database("cp-server").Collection("posts")
	ctx := context.Background()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.New(os.Stdout, "primitive.ObjectIDFromHex ERROR: "+err.Error()+"\t"+"", 1)
		return nil
	}

	err = coll.FindOne(ctx, bson.M{"_id": objectId}).Decode(&post)
	if err != nil {
		log.New(os.Stdout, "coll.FindOne ERROR: "+err.Error()+"\t"+"", 1)
		return nil
	}

	return &post
}
