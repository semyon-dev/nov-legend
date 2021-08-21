package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nov-legend/app/db"
	"nov-legend/app/session"
)

func Auth(c *gin.Context) {

	jsonInput := struct {
		Phone string `json:"phone"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	user, exist := db.GetUserByPhone(jsonInput.Phone)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	token, refreshToken, err := session.Create(user.Id.Hex())
	if err != nil {
		fmt.Println("Error in generating JWT: " + err.Error())
	}

	err = db.UpdateRefreshToken(user.Id, refreshToken)
	if err != nil {
		log.Println("can't update token: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "ok",
		"token":        token,
		"refreshToken": refreshToken,
	})
}
