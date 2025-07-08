package repository

import "github.com/efectn/fiber-boilerplate/internal/database"

type ArticleRepository struct {
	DB *database.Database
}

type Article struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (Article) TableName() string {
	return "articles"
}

type IArticleRepository interface {
	GetArticles() ([]*Article, error)
	GetArticleByID(id uint) (*Article, error)
	CreateArticle(title, content string) (*Article, error)
	UpdateArticle(id uint, title, content string) (*Article, error)
	DeleteArticle(id uint) error
}

func NewArticleRepository(database *database.Database) *ArticleRepository {
	return &ArticleRepository{
		DB: database,
	}
}

func (s *ArticleRepository) GetArticles() ([]*Article, error) {
	var articles []*Article
	err := s.DB.Gorm.Order("id asc").Find(&articles).Error
	return articles, err
}

func (s *ArticleRepository) GetArticleByID(id uint) (*Article, error) {
	var article Article
	err := s.DB.Gorm.Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (s *ArticleRepository) CreateArticle(title, content string) (*Article, error) {
	article := &Article{
		Title:   title,
		Content: content,
	}
	err := s.DB.Gorm.Create(article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleRepository) UpdateArticle(id uint, title, content string) (*Article, error) {
	var article Article
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

func (s *ArticleRepository) DeleteArticle(id uint) error {
	return s.DB.Gorm.Delete(&Article{}, id).Error
}
