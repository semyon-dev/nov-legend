package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	AccessSecret string
	MongoUrl     string
	Port         string
	FolderId     string
	IAMTOKEN     string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can't load from file: " + err.Error())
	}
	AccessSecret = os.Getenv("ACCESS_SECRET")
	MongoUrl = os.Getenv("MONGO_URL")
	Port = os.Getenv("PORT")
	FolderId = os.Getenv("FOLDER_ID")
	IAMTOKEN = os.Getenv("IAM_TOKEN")
}
