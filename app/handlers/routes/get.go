package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
)

func GetRoutes(c *gin.Context) {
	routes := db.GetRoutes()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"routes":  routes,
	})
}

func SelectRoutes(c *gin.Context) {
	input := struct {
		Type     string `json:"type"`
		Company  string `json:"company"`
		FoodNeed bool   `json:"foodNeed"`
		Duration uint   `json:"duration"`
	}{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid params",
		})
	}
	routes := db.GetRoutesByType(input.Type)
	var routesReply []model.Route
	for _, v := range routes {
		if v.Duration == input.Duration {
			routesReply = append(routesReply, v)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"routes":  routesReply,
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
