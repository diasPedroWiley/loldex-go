package main

import (
	"fmt"
	"log"

	"github.com/diasPedroWiley/loldex-go/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func inital() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	fmt.Println("Hello")
	r.GET("/champions", controller.GetChampions)
	r.Run()
}
