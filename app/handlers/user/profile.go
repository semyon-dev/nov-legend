package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/model"
	"nov-legend/app/session"
)

func Update(c *gin.Context) {
	_, isValid := session.ParseBearer(c)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}
	if user.Id.IsZero() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user id can't be empty",
		})
		return
	}
	err := db.UpdateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func GetUser(c *gin.Context) {
	id, isValid := session.ParseBearer(c)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}
	user, ok := db.GetUserByIdString(id)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	user.Level = user.Exp / 1000

	if len(user.AchievementsIds) != 0 {
		user.Achievements = db.GetAchievementsByIds(user.AchievementsIds)
	} else {
		user.Achievements = []model.Achievement{}
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}
	user, ok := db.GetUserById(objectId)
	if ok {
		user.Level = user.Exp / 1000
		c.JSON(http.StatusOK, gin.H{
			"user":    user,
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "user not found",
	})
}
