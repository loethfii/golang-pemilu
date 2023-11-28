package domain

import (
	"context"
	"time"
)

type Article struct {
	ID          int64     `gorm:"type:bigint; primaryKey; autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(100)" json:"title"`
	Author      string    `gorm:"type:varchar(100)" json:"author"`
	Image       string    `gorm:"type:varchar(300)" json:"image"`
	Description string    `json:"description"`
	PostedAt    time.Time `gorm:"type:timestamp with time zone; default:CURRENT_TIMESTAMP" json:"posted_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type ArticleResponses struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	PostedAt    time.Time `json:"posted_at"`
}

type ArticleResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	PostedAt    time.Time `json:"posted_at"`
}

type ArticleResponseUpdate struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	PostedAt    time.Time `json:"posted_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type ArticleRepository interface {
	Fetch(ctx context.Context) ([]Article, error)
	GetByID(ctx context.Context, id int64) (Article, error)
	Store(ctx context.Context, article Article) (Article, error)
	Update(ctx context.Context, id int64, article Article) (Article, error)
	Delete(ctx context.Context, id int64) error
}

type ArticleUseCase interface {
	Fetch() ([]ArticleResponses, error)
	GetByID(id int64) (ArticleResponse, error)
	Store(article Article) (Article, error)
	Update(id int64, article Article) (ArticleResponseUpdate, error)
	Delete(id int64) error
}
