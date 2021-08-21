package routes

import (
	"fmt"
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
	route, isExist := db.GetRouteByIDString(id)
	if !isExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "route not found",
			"route":   route,
		})
		return
	}
	for i, v := range route.Steps {
		fmt.Println(v.PointId)
		point, isExist := db.GetPointByID(v.PointId)
		if !isExist {
			continue
		}
		route.Steps[i].Point = point
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"route":   route,
	})
}
