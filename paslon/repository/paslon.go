package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"luthfi/pemilu/domain"
)

type paslonRepository struct {
	*gorm.DB
}

func NewPaslonRepository(db *gorm.DB) domain.PaslonRepository {
	return &paslonRepository{db}
}

func (r *paslonRepository) Fetch(ctx context.Context) ([]domain.Paslon, error) {
	var paslon []domain.Paslon
	err := r.DB.WithContext(ctx).Order("id desc").Find(&paslon).Error
	if err != nil {
		return nil, err
	}
	
	return paslon, nil
}

func (r *paslonRepository) GetByID(ctx context.Context, id int64) (domain.Paslon, error) {
	var paslon domain.Paslon
	res := r.DB.WithContext(ctx).First(&paslon, id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return domain.Paslon{}, errors.New("paslon not found")
		}
		
		return domain.Paslon{}, res.Error
	}
	
	return paslon, nil
}

func (r *paslonRepository) Store(ctx context.Context, paslon domain.Paslon) (domain.Paslon, error) {
	err := r.DB.WithContext(ctx).Create(&paslon).Error
	if err != nil {
		return domain.Paslon{}, err
	}
	
	return paslon, nil
}

func (r *paslonRepository) Update(ctx context.Context, id int64, paslon domain.Paslon) (domain.Paslon, error) {
	res := r.DB.WithContext(ctx).Where("id = ?", id).Updates(&paslon)
	if res.RowsAffected == 0 {
		return domain.Paslon{}, errors.New("paslon not found")
	}
	
	return paslon, nil
}

func (r *paslonRepository) Delete(ctx context.Context, id int64) error {
	var paslon domain.Paslon
	res := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&paslon)
	if res.RowsAffected == 0 {
		return errors.New("paslon not found")
	}
	
	return nil
}
