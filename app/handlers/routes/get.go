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
		"points":  routes,
	})
}
