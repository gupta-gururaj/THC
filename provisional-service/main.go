package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"provisional-service/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading env variables", err)
	}

	r := gin.Default()
	r.POST("/ping", controller.ABC)
	r.Run(os.Getenv("ADDRESS") + ":" + os.Getenv("HTTPPORT"))
}
