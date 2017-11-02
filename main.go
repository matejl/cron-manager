package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matejl/cron-manager/controller"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "static")

	router.GET("/index", controller.Index)

	router.Run(":8080")
}