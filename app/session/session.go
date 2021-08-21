package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"math/rand"
	"net/http"
	"nov-legend/app/config"
	"strings"
	"time"
)

func Create(id string) (token string, refreshToken string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(config.AccessSecret))
	if err != nil {
		return "", "", err
	}
	return token, NewRefreshToken(), err
}

type MyCustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func ParseToken(tokenString string) (id string) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AccessSecret), nil
	})

	if token == nil {
		log.Println("empty token")
		return
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		id = claims.Id
	} else {
		fmt.Println(err)
	}
	return
}

func NewRefreshToken() string {
	b := make([]byte, 32)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	if _, err := r.Read(b); err != nil {
		return ""
	}

	return fmt.Sprintf("%x", b)
}

func ParseBearer(c *gin.Context) (id string, isValid bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	id = ParseToken(headerParts[1])
	isValid = true
	return
}
