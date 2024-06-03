package main

import (
	"github.com/Cautery/go-gin-demo/handlers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.ShowIndexPage)

	router.Run()
}
