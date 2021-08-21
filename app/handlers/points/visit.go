package points

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
	"nov-legend/app/session"
	"nov-legend/app/util"
)

func Visit(c *gin.Context) {

	userId, isValid := session.ParseBearer(c)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}

	jsonInput := struct {
		PointId  string            `json:"pointId"`
		Location model.Coordinates `json:"location"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		log.Println(err)
		return
	}

	point, isExist := db.GetPointByIDString(jsonInput.PointId)
	if !isExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "point not found",
			"check":   false,
			"exp":     0,
		})
		return
	}

	km := util.Distance(point.Coordinates.Lat, point.Coordinates.Lng, jsonInput.Location.Lat, jsonInput.Location.Lng)
	if km < 10000.00 {
		db.AddExpToUser(userId, 500)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"check":   true,
			"exp":     500,
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "ok",
		"check":   false,
		"exp":     0,
	})
}
