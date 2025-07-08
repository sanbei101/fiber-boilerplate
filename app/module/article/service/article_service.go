package service

import (
	"github.com/efectn/fiber-boilerplate/app/module/article/repository"
)

type ArticleService struct {
	Repo *repository.ArticleRepository
}

type IArticleService interface {
	GetArticles() ([]*repository.Article, error)
	GetArticleByID(id uint) (*repository.Article, error)
	CreateArticle(id uint, title string, content string) (*repository.Article, error)
	UpdateArticle(id uint, title string, content string) (*repository.Article, error)
	DeleteArticle(id uint) error
}

func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		Repo: repo,
	}
}
func (s *ArticleService) GetArticles() ([]*repository.Article, error) {
	return s.Repo.GetArticles()
}

func (s *ArticleService) GetArticleByID(id uint) (*repository.Article, error) {
	return s.Repo.GetArticleByID(id)
}

func (s *ArticleService) CreateArticle(title string, content string) (*repository.Article, error) {
	return s.Repo.CreateArticle(title, content)
}
func (s *ArticleService) UpdateArticle(id uint, title string, content string) (*repository.Article, error) {
	return s.Repo.UpdateArticle(id, title, content)
}

func (s *ArticleService) DeleteArticle(id uint) error {
	return s.Repo.DeleteArticle(id)
}
