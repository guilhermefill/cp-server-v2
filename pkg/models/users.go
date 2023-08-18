package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Avatar    string             `bson:"avatar"`
	Email     string             `bson:"email"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	UserType  string             `bson:"userType"`
}

func GetUsers() *[]User {
	var users []User
	coll := Client.Database("cp-server").Collection("users")
	findOptions := options.Find()

	ctx := context.Background()

	cursor, err := coll.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var result User
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal("cursor.Decode ERROR:", err)
		}
		users = append(users, result)
	}

	defer cursor.Close(ctx)

	return &users
}
