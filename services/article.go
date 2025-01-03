package services

import "github.com/cleanupDev/blog/internal"

type ArticleService struct {
	articles []internal.Article
}

func NewArticleService() *ArticleService {
	articles := []internal.Article{
		{ID: 1, Title: "Hello World", Summary: "My first post!", Content: "Full content..."},
		{ID: 2, Title: "Another Post", Summary: "Check it out!", Content: "More details..."},
	}
	return &ArticleService{articles: articles}
}

func (s *ArticleService) GetAllArticles() []internal.Article {
	return s.articles
}
