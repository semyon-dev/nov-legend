package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
	"nov-legend/app/session"
)

func CreateRoute(c *gin.Context) {
	id, isValid := session.ParseBearer(c)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}
	var route model.Route
	err := c.ShouldBindJSON(route)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
		})
		return
	}

	route.Id = primitive.NewObjectID()
	ojb, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}
	route.AuthorId = ojb
	route.Comments = []model.Comment{}

	err = db.Insert("routes", route)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "internal server error " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"routeId": route.Id.Hex(),
	})
}
