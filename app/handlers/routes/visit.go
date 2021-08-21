package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
)

func Visit(c *gin.Context) {
	jsonInput := struct {
		RouteId  string            `json:"routeId"`
		Location model.Coordinates `json:"location"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		log.Println(err)
		return
	}

	_, isExist := db.GetRouteByIDString(jsonInput.RouteId)
	if !isExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "route not found",
			"check":   false,
			"exp":     0,
		})
		return
	}

	//km := util.Distance(point.Coordinate.Lat, point.Coordinate.Lng, jsonInput.Location.Lat, jsonInput.Location.Lng)
	//if km < 10000.00 {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"check":   true,
		"exp":     500,
	})
	//	return
	//}
	//c.AbortWithStatusJSON(http.StatusOK, gin.H{
	//	"message": "ok",
	//	"check":   false,
	//	"exp":     0,
	//})
}
