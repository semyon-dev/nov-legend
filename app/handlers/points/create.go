package points

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
	"nov-legend/app/session"
	"nov-legend/app/util"
)

func CreatePoint(c *gin.Context) {
	id, isValid := session.ParseBearer(c)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}
	var point model.Point
	if err := c.ShouldBindJSON(&point); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		log.Println(err)
		return
	}
	point.Id = primitive.NewObjectID()
	point.DescriptionEN = util.Translate("en", point.Description)
	point.Comments = []model.Comment{}

	point.Id = primitive.NewObjectID()
	ojb, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
	}
	point.AuthorId = ojb

	err = db.Insert("points", point)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "ok",
	})
}
