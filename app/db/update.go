package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"nov-legend/app/model"
)

func UpdateUser(user model.User) (err error) {
	filter := bson.M{"_id": user.Id}

	var thisUser model.User
	result := db.Collection("users").FindOne(context.Background(), filter)
	err = result.Decode(&thisUser)
	if err != nil {
		log.Println(err)
		return err
	}

	result = db.Collection("users").FindOneAndReplace(context.Background(), filter, user)
	if result.Err() != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			log.Println("no docs")
		}
		return
	}
	return nil
}

func AddExpToUser(userId string, exp uint) {

}

func UpdateRefreshToken(userId primitive.ObjectID, token string) error {
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"refreshToken": token}}
	_, err := db.Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return err
}
