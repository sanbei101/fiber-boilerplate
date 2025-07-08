package model

type Article struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (Article) TableName() string {
	return "articles"
}
