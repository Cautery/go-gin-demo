package handlers

import (
	"net/http"
	"strconv"

	"github.com/Cautery/go-gin-demo/models"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi((c.Param("article_id")))
	if err != nil { // invalid article ID
		c.AbortWithStatus(http.StatusNotFound)
	}

	article, err := models.GetArticleByID(articleID)
	if err != nil { // article not found
		c.AbortWithError(http.StatusNotFound, err)
	}

	// no errors, then render template
	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title": article.Title,
			"payload": article,
		},
	)
}