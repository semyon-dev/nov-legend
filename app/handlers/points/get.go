package points

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
	"nov-legend/app/util"
)

func GetPoints(c *gin.Context) {
	points := db.GetPoints()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"points":  points,
	})
}

func GetPoint(c *gin.Context) {
	id := c.Param("id")
	obj, _ := primitive.ObjectIDFromHex(id)
	point, _ := db.GetPointByID(obj)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"point":   point,
	})
}

func GetPointsByText(c *gin.Context) {
	name := c.Param("name")
	points := db.FindPointsByText(name)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"points":   points,
	})
}

func GetDistance(c *gin.Context) {

	json := struct {
		CurrentLocation model.Coordinate `json:"currentLocation"`
		Points          []string         `json:"points"`
	}{}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	var reply = map[string]float64{}

	var pointsP []primitive.ObjectID
	for _, v := range json.Points {
		ob, _ := primitive.ObjectIDFromHex(v)
		pointsP = append(pointsP, ob)
	}
	points := db.GetPointsByIds(pointsP)

	for _, v := range points {
		reply[v.Id.Hex()] = util.Distance(json.CurrentLocation.Lat, json.CurrentLocation.Lng, v.Coordinate.Lat, v.Coordinate.Lng)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"reply":   reply,
	})
}
