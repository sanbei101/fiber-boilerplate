package service

import (
	"github.com/efectn/fiber-boilerplate/app/module/article/model"
	"github.com/efectn/fiber-boilerplate/internal/database"
)

type ArticleService struct {
	DB *database.Database
}

type IArticleService interface {
	GetArticles() ([]*model.Article, error)
	GetArticleByID(id uint) (*model.Article, error)
	CreateArticle(title string, content string) (*model.Article, error)
	UpdateArticle(id uint, title string, content string) (*model.Article, error)
	DeleteArticle(id uint) error
}

func NewArticleService(db *database.Database) *ArticleService {
	return &ArticleService{
		DB: db,
	}
}

func (s *ArticleService) GetArticles() ([]*model.Article, error) {
	var articles []*model.Article
	err := s.DB.Gorm.Order("id asc").Find(&articles).Error
	return articles, err
}

func (s *ArticleService) GetArticleByID(id uint) (*model.Article, error) {
	var article model.Article
	err := s.DB.Gorm.Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (s *ArticleService) CreateArticle(title string, content string) (*model.Article, error) {
	article := &model.Article{
		Title:   title,
		Content: content,
	}
	err := s.DB.Gorm.Create(article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) UpdateArticle(id uint, title string, content string) (*model.Article, error) {
	var article model.Article
	err := s.DB.Gorm.Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}
	article.Title = title
	article.Content = content
	err = s.DB.Gorm.Save(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (s *ArticleService) DeleteArticle(id uint) error {
	return s.DB.Gorm.Delete(&model.Article{}, id).Error
}
