package main

import (
	"github.com/gin-gonic/gin"

	"main/controllers" // new
	"main/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
