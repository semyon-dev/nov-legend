package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nov-legend/app/db"
)

func GetRoutes(c *gin.Context) {
	routes := db.GetRoutes()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"routes":  routes,
	})
}

func GetRoute(c *gin.Context) {
	id := c.Param("id")
	route, _ := db.GetRouteByID(id)
	for i, v := range route.Steps {
		point, _ := db.GetPointByID(v.PointId)
		route.Steps[i].Coordinates = point.Coordinates
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"route":   route,
	})
}
