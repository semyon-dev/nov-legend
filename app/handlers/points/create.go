package points

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
	"nov-legend/app/util"
)

func CreatePoint(c *gin.Context) {
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
	err := db.Insert("points", point)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
