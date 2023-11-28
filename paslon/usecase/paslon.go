package usecase

import (
	"context"
	"luthfi/pemilu/domain"
)

type paslonUseCase struct {
	domain.PaslonRepository
}

func NewPaslonUseCase(paslonRepository domain.PaslonRepository) domain.PaslonUseCase {
	return &paslonUseCase{paslonRepository}
}

func (u *paslonUseCase) Fetch() ([]domain.PaslonResponses, error) {
	ctx := context.Background()
	paslons, err := u.PaslonRepository.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	
	var paslonsResponses []domain.PaslonResponses
	for _, paslon := range paslons {
		paslonResponses := domain.PaslonResponses{
			ID:            paslon.ID,
			Name:          paslon.Name,
			SerialNumber:  paslon.SerialNumber,
			VisionMission: paslon.VisionMission,
			Image:         paslon.Image,
		}
		paslonsResponses = append(paslonsResponses, paslonResponses)
	}
	
	return paslonsResponses, nil
}

func (u *paslonUseCase) GetByID(id int64) (domain.PaslonResponse, error) {
	ctx := context.Background()
	
	paslon, err := u.PaslonRepository.GetByID(ctx, id)
	if err != nil {
		return domain.PaslonResponse{}, err
	}
	
	paslonResponse := domain.PaslonResponse{
		ID:            paslon.ID,
		Name:          paslon.Name,
		SerialNumber:  paslon.SerialNumber,
		VisionMission: paslon.VisionMission,
		Image:         paslon.Image,
	}
	
	return paslonResponse, nil
}

func (u *paslonUseCase) Store(paslon domain.Paslon) (domain.Paslon, error) {
	ctx := context.Background()
	paslon, err := u.PaslonRepository.Store(ctx, paslon)
	if err != nil {
		return domain.Paslon{}, err
	}
	
	return paslon, nil
}

func (u *paslonUseCase) Update(id int64, paslon domain.Paslon) (domain.Paslon, error) {
	ctx := context.Background()
	paslon, err := u.PaslonRepository.Update(ctx, id, paslon)
	if err != nil {
		return domain.Paslon{}, err
	}
	
	return paslon, nil
}

func (u *paslonUseCase) Delete(id int64) error {
	ctx := context.Background()
	err := u.PaslonRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	
	return nil
}
