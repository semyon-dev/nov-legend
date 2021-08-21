package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"nov-legend/app/model"
)

func GetPointByID(id interface{}) (point model.Point, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection("points").FindOne(context.Background(), filter).Decode(&point)
	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			return model.Point{}, false
		}
		return
	}
	return point, true
}

func FindPointsByText(text string) (points []model.Point) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	//opts := options.Find().SetLimit(50).SetMaxTime(time.Second * 3)
	cursor, err := db.Collection("points").Find(context.Background(), filter, nil)
	if err != nil {
		log.Println(err)
	}
	// выделяем память заранее
	points = make([]model.Point, 0, 50)
	err = cursor.All(context.Background(), &points)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetRouteByID(id interface{}) (point model.Route, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection("routes").FindOne(context.Background(), filter).Decode(&point)
	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			return model.Route{}, false
		}
		return
	}
	return point, true
}

func GetPoints() (points []model.Point) {
	cursor, err := db.Collection("points").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &points); err != nil {
		log.Println(err)
	}
	return
}

func GetRoutes() (points []model.Route) {
	cursor, err := db.Collection("routes").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &points); err != nil {
		log.Println(err)
	}
	return
}

func GetPointsByIds(ids []primitive.ObjectID) (points []model.Point) {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	cursor, err := db.Collection("points").Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &points); err != nil {
		log.Println(err)
	}
	return
}

func GetUsersByIds(ids []primitive.ObjectID) (users []model.User) {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	cursor, err := db.Collection("users").Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &users); err != nil {
		log.Println(err)
	}
	return
}

func GetAchievementsByIds(ids []primitive.ObjectID) (achievements []model.Achievement) {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	cursor, err := db.Collection("achievements").Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	if err = cursor.All(context.Background(), &achievements); err != nil {
		log.Println(err)
	}
	return
}

func GetUserByPhone(phone string) (user model.User, isExist bool) {
	filter := bson.M{"phone": phone}
	err := db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		return
	}
	return user, true
}

func GetUserById(id primitive.ObjectID) (user model.User, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, false
		}
		return
	}
	return user, true
}

func GetUserByIdString(id string) (user model.User, isExist bool) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	return GetUserById(objId)
}
