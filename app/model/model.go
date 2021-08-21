package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id              primitive.ObjectID            `json:"id" bson:"_id"`
	RefreshToken    string                        `json:"-" bson:"refreshToken"`
	Name            string                        `json:"name" bson:"name"`
	Email           string                        `json:"email" bson:"email"`
	Phone           string                        `json:"phone" bson:"phone"`
	PhotoURL        string                        `json:"photoURL" bson:"photoURL"`
	Exp             int                           `json:"exp" bson:"exp"`
	Level           int                           `json:"level" bson:"-"`
	CompletedRoutes []primitive.ObjectID          `json:"completedStories" bson:"completedStories"`
	Achievements    map[string]primitive.DateTime `json:"achievements" bson:"achievements"`
	CreatedAt       time.Time                     `json:"-" bson:"createdAt"`
	Role            string                        `json:"role" bson:"role"` // user or moderator
}

type Route struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Step []struct {
		Name        string             `json:"name" bson:"name"`
		Description string             `json:"description" bson:"description"`
		PhotoURL    string             `json:"photoURL" bson:"photoURL"`
		PointId     primitive.ObjectID `json:"pointId" bson:"pointId"`
	} `json:"step" bson:"step"`
	Duration string    `json:"duration" bson:"duration"`
	Tags     []string  `json:"tags" bson:"tags"`
	Exp      int       `json:"exp" bson:"exp"`
	Comments []Comment `json:"comments" bson:"comments"`
	Likes    uint      `json:"likes" bson:"likes"`
}

type Comment struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Text     string             `json:"text" bson:"text"`
	AuthorId primitive.ObjectID `json:"authorId" bson:"authorId"`
	Date     primitive.DateTime `json:"date" bson:"date"`
	Likes    uint               `json:"likes" bson:"likes"`
}

type Achievements struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Exp         int                `json:"exp" bson:"exp"`
	Icon        string             `json:"icon" bson:"icon"`
}

type Coordinate struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

type Point struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	AuthorId      primitive.ObjectID `json:"authorId" bson:"authorId"`
	PhotoURL      string             `json:"photoURL" bson:"photoURL"`
	HowToGet      string             `json:"howToGet" bson:"howToGet"`
	Name          string             `json:"name" bson:"name"`
	Type          string             `json:"type" bson:"type"`
	Tags          []string           `json:"tags" bson:"tags"`
	Coordinate    Coordinate         `json:"coordinate" bson:"coordinate"`
	Description   string             `json:"description" bson:"description"`
	DescriptionEN string             `json:"descriptionEN" bson:"descriptionEN"`
	Website       string             `json:"website" bson:"website"`
	Comments      []Comment          `json:"comments" bson:"comments"`
	Likes         uint               `json:"likes" bson:"likes"`
}