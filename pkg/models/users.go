package models

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Email     string             `bson:"email" json:"email"`
	FirstName string             `bson:"firstName" json:"first_name"`
	LastName  string             `bson:"lastName" json:"last_name"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password"`
	UserType  string             `bson:"userType" json:"user_type"`
}

func GetUsers() *[]User {
	var users []User
	coll := Client.Database("cp-server").Collection("users")
	findOptions := options.Find()

	ctx := context.Background()

	cursor, err := coll.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.New(os.Stdout, "coll.Find ERROR:"+err.Error()+"\t"+"", 1)
		return nil
	}

	for cursor.Next(ctx) {
		var result User
		err := cursor.Decode(&result)
		if err != nil {
			log.New(os.Stdout, "cursor.Decode ERROR:"+err.Error()+"\t"+"", 1)
			return nil
		}
		users = append(users, result)
	}

	defer cursor.Close(ctx)

	return &users
}
