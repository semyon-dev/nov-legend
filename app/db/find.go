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

func GetPointByID(id primitive.ObjectID) (point model.Point, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection("points").FindOne(context.Background(), filter).Decode(&point)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return model.Point{}, false
		}
		return
	}
	return point, true
}

func GetRoutesByType(typ string) (routes []model.Route) {
	filter := bson.M{"type": typ}
	cursor, err := db.Collection("routes").Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	// выделяем память заранее
	routes = make([]model.Route, 0, 50)
	err = cursor.All(context.Background(), &routes)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetPointByIDString(id string) (point model.Point, isExist bool) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	return GetPointByID(objId)
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

func GetRouteByID(id primitive.ObjectID) (route model.Route, isExist bool) {
	filter := bson.M{"_id": id}
	err := db.Collection("routes").FindOne(context.Background(), filter).Decode(&route)
	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			return model.Route{}, false
		}
		return
	}
	return route, true
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

func GetRouteByIDString(id string) (route model.Route, isExist bool) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	return GetRouteByID(objId)
}
