package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"luthfi/pemilu/domain"
)

type articleRepository struct {
	*gorm.DB
}

func NewArticleRepository(db *gorm.DB) domain.ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) Fetch(ctx context.Context) ([]domain.Article, error) {
	var articles []domain.Article
	err := r.DB.WithContext(ctx).Order("id desc").Find(&articles).Error
	if err != nil {
		return nil, err
	}
	
	return articles, nil
}

func (r *articleRepository) GetByID(ctx context.Context, id int64) (domain.Article, error) {
	var article domain.Article
	res := r.DB.WithContext(ctx).First(&article, id)
	
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return domain.Article{}, errors.New("article not found")
		}
		return domain.Article{}, res.Error
	}
	
	return article, nil
}

func (r *articleRepository) Store(ctx context.Context, article domain.Article) (domain.Article, error) {
	err := r.DB.WithContext(ctx).Create(&article).Error
	if err != nil {
		return domain.Article{}, err
	}
	
	return article, nil
}

func (r *articleRepository) Update(ctx context.Context, id int64, article domain.Article) (domain.Article, error) {
	res := r.DB.WithContext(ctx).Where("id = ?", id).Updates(&article)
	if res.RowsAffected == 0 {
		return domain.Article{}, errors.New("article not found")
	}
	
	return article, nil
}

func (r *articleRepository) Delete(ctx context.Context, id int64) error {
	var article domain.Article
	res := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&article)
	if res.RowsAffected == 0 {
		return errors.New("article not found")
	}
	
	return nil
}
