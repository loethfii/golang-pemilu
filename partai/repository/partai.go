package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"luthfi/pemilu/domain"
)

type partaiRepository struct {
	*gorm.DB
}

func NewPartaiRepository(db *gorm.DB) domain.PartaiRepository {
	return &partaiRepository{db}
}

func (r *partaiRepository) Fetch(ctx context.Context) ([]domain.Partai, error) {
	var partais []domain.Partai
	
	err := r.DB.WithContext(ctx).Order("id desc").Find(&partais).Error
	if err != nil {
		return nil, err
	}
	
	return partais, nil
}

func (r *partaiRepository) GetByID(ctx context.Context, id int64) (domain.Partai, error) {
	var partai domain.Partai
	res := r.DB.WithContext(ctx).First(&partai, id)
	
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return domain.Partai{}, errors.New("partai not found")
		}
		
		return domain.Partai{}, res.Error
	}
	
	return partai, nil
}

func (r *partaiRepository) Store(ctx context.Context, partai domain.Partai) (domain.Partai, error) {
	err := r.DB.WithContext(ctx).Create(&partai).Error
	if err != nil {
		return domain.Partai{}, err
	}
	
	return partai, nil
}

func (r *partaiRepository) Update(ctx context.Context, id int64, partai domain.Partai) (domain.Partai, error) {
	res := r.DB.WithContext(ctx).Where("id = ?", id).Updates(&partai)
	if res.RowsAffected == 0 {
		return domain.Partai{}, errors.New("partai not found")
	}
	
	return partai, nil
	
}

func (r *partaiRepository) Delete(ctx context.Context, id int64) error {
	var partai domain.Partai
	res := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&partai)
	if res.RowsAffected == 0 {
		return errors.New("partai not found")
	}
	
	return nil
}
