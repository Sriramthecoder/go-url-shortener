package main

import (
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:code", handlers.ResolveURL)

	r.Run(":8080")
}
