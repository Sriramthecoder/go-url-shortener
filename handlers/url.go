package handlers

import (
	"net/http"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

type URLRequest struct {
	URL string `json:"url"`
}

func ShortenURL(c *gin.Context) {
	var req URLRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	code, _ := shortid.Generate()
	store.Save(code, req.URL)

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + code,
	})
}

func ResolveURL(c *gin.Context) {
	code := c.Param("code")
	url := store.Get(code)

	if url == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, url)
}
