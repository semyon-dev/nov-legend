package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nov-legend/app/session"
)

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

func Refresh(c *gin.Context) {
	var inp refreshInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	res := session.NewRefreshToken()
	//session.Create()

	c.JSON(http.StatusOK, gin.H{
		"AccessToken":  res,
		"RefreshToken": res,
	})
}
