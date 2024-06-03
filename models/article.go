package models

import (
	"errors"
	"os"
)

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

var ArticleList = []Article{
	{ID: 1, Title: "Metamorphosis", Author: "Franz Kafka", Content: readArticle("metamorphosis.txt")},
	{ID: 2, Title: "The War of the Worlds", Author: "H. G. Wells", Content: readArticle("war-of-the-worlds.txt")},
}

func readArticle(file string) string {
	text, err := os.ReadFile("resources/" + file)
	if err != nil {
		return "Unable to read file."
	}
	return string(text)
}

func GetAllArticles() []Article {
	return ArticleList
}

func GetArticleByID(id int) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("Article not found")
}
