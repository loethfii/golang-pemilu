package usecase

import (
	"context"
	"fmt"
	"luthfi/pemilu/domain"
	"time"
)

type articleUseCase struct {
	domain.ArticleRepository
}

func NewArticleUseCase(ar domain.ArticleRepository) domain.ArticleUseCase {
	return &articleUseCase{ar}
}

func (u *articleUseCase) Fetch() ([]domain.ArticleResponses, error) {
	var ctx = context.Background()
	deadline := time.Now().Add(2 * time.Second)
	ctxWithDeadline, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	articles, err := u.ArticleRepository.Fetch(ctxWithDeadline)
	if err != nil {
		return nil, err
	}
	
	var errTimeOut = make(chan string)
	
	go func() {
		select {
		case <-ctxWithDeadline.Done():
			errTimeOut <- "Find All article sengaja gua buat buffering 2 detik, klo mau matiin bisa dihapus deadlinenya "
		}
	}()
	
	fmt.Println(<-errTimeOut)
	
	var articleResponses []domain.ArticleResponses
	for _, article := range articles {
		articleResponses = append(articleResponses, domain.ArticleResponses{
			ID:          article.ID,
			Title:       article.Title,
			Author:      article.Author,
			Image:       article.Image,
			Description: article.Description,
			PostedAt:    article.PostedAt,
		})
	}
	
	return articleResponses, nil
}

func (u *articleUseCase) GetByID(id int64) (domain.ArticleResponse, error) {
	var ctx = context.Background()
	article, err := u.ArticleRepository.GetByID(ctx, id)
	var articleResponse = domain.ArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		Author:      article.Author,
		Image:       article.Image,
		Description: article.Description,
		PostedAt:    article.PostedAt,
	}
	if err != nil {
		return articleResponse, err
	}
	
	return articleResponse, nil
}

func (u *articleUseCase) Store(article domain.Article) (domain.Article, error) {
	var ctx = context.Background()
	deadline := time.Now().Add(10 * time.Minute)
	ctxWithDeadline, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	article, err := u.ArticleRepository.Store(ctxWithDeadline, article)
	if err != nil {
		return domain.Article{}, err
	}
	return article, nil
}

func (u *articleUseCase) Update(id int64, article domain.Article) (domain.ArticleResponseUpdate, error) {
	ctx := context.Background()
	article, err := u.ArticleRepository.Update(ctx, id, article)
	if err != nil {
		return domain.ArticleResponseUpdate{}, err
	}
	
	articleID, _ := u.ArticleRepository.GetByID(ctx, id)
	
	var articleResponseUpdate = domain.ArticleResponseUpdate{
		ID:          id,
		Title:       article.Title,
		Author:      article.Author,
		Image:       article.Image,
		Description: article.Description,
		PostedAt:    articleID.PostedAt,
		UpdatedAt:   article.UpdatedAt,
		CreatedAt:   articleID.CreatedAt,
	}
	
	return articleResponseUpdate, nil
}

func (u *articleUseCase) Delete(id int64) error {
	ctx := context.Background()
	err := u.ArticleRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	
	return nil
}
