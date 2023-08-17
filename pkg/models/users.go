package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	_id        primitive.ObjectID `bson:"_id"`
	avatar     string             `bson:"avatar"`
	email      string             `bson:"email"`
	favourites []string           `bson:"favourite,omitempty"`
	firstName  string             `bson:"firstName"`
	lastName   string             `bson:"lastName"`
	username   string             `bson:"username"`
	password   string             `bson:"password"`
	userType   string             `bson:"userType"`
}

func GetUsers() *[]User {
	var users []User
	coll := Client.Database("cp-server").Collection("users")
	findOptions := options.Find()
	cursor, err := coll.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var result bson.D
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range result {
			fmt.Println(v.Key, v.Value)
		}
		// do something with result....
	}
	// if err = cursor.All(context.TODO(), &users); err != nil {
	// 	panic(err)
	// }

	defer cursor.Close(context.TODO())

	for _, user := range users {
		fmt.Println(user._id, user.email, user.username, user.password, user.avatar, user.userType, user.firstName, user.lastName, user.favourites)
	}
	return &users
}
