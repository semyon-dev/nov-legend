package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
	"nov-legend/app/session"
	"time"
)

func Register(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	if user.PhotoURL == "" {
		user.PhotoURL = "https://cdn.onlinewebfonts.com/svg/img_341152.png"
	}

	userToken, refreshToken, err := session.Create(user.Id.Hex())
	if err != nil {
		fmt.Println("Error in generating JWT: " + err.Error())
	}

	user.RefreshToken = refreshToken

	user.AchievementsIds = []primitive.ObjectID{}

	err = db.Insert("users", user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "ok",
		"token":        userToken,
		"refreshToken": refreshToken,
	})
}
