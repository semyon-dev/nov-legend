package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"nov-legend/app/config"
	"nov-legend/app/db"
	"nov-legend/app/handlers/points"
	"nov-legend/app/handlers/routes"
	"nov-legend/app/handlers/user"
)

func main() {

	config.Load()
	db.Connect()

	app := gin.Default()
	app.Use(cors.Default())

	gin.SetMode(gin.DebugMode)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})

	// users
	app.POST("/auth", user.Auth)
	app.POST("/user/refresh", user.Refresh) // TODO
	//app.GET("/user/token", user) // TODO
	app.POST("/user", user.Register)
	app.PUT("/user", user.Update)
	app.GET("/user", user.GetUser)
	app.GET("/user/:id", user.GetUserById)

	// points
	app.GET("/points", points.GetPoints)             // get points
	app.GET("/point/:id", points.GetPoint)           // get point by ID
	app.GET("/point/:name", points.GetPointsByText)  // get point by TEXT
	app.POST("/point", points.CreatePoint)           // create
	app.POST("/point/visit", points.Visit)           // visit point
	app.POST("/points/distance", points.GetDistance) // get distance of point(s)

	// routes
	app.GET("/routes", routes.GetRoutes)   // get routes
	app.POST("/route/visit", routes.Visit) // visit route

	err := app.Run("0.0.0.0:" + config.Port)
	if err != nil {
		fmt.Println("Error in launching backend: " + err.Error())
	}
}
