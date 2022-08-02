package main

import (
	"github.com/gin-gonic/gin"

	"main/models" // new
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.Run()
}
