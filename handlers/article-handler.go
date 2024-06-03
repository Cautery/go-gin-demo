package handlers

import (
	"net/http"
	"strconv"

	"github.com/Cautery/go-gin-demo/models"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	a := models.GetAllArticles()
	render(c, gin.H{"title": "Home Page", "payload": a}, "index.html")
}

func GetArticle(c *gin.Context) {
	aID, err := strconv.Atoi((c.Param("article_id")))
	if err != nil { // invalid article ID
		c.AbortWithStatus(http.StatusNotFound)
	}

	a, err := models.GetArticleByID(aID)
	if err != nil { // article not found
		c.AbortWithError(http.StatusNotFound, err)
	}

	// no errors, render template
	render(c, gin.H{"title": a.Title, "payload": a}, "article.html")
}

func render(c *gin.Context, data gin.H, template string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, template, data)
	}
}
