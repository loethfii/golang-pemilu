package usecase

import (
	"context"
	"luthfi/pemilu/domain"
)

type partaiUseCase struct {
	partaiRepository domain.PartaiRepository
}

func NewPartaiUseCase(partaiRepository domain.PartaiRepository) domain.PartaiUseCase {
	return &partaiUseCase{partaiRepository}
}

func (u *partaiUseCase) Fetch() ([]domain.PartaiResponses, error) {
	ctx := context.Background()
	partais, err := u.partaiRepository.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	
	var partaisResponses []domain.PartaiResponses
	for _, partai := range partais {
		partaisResponses = append(partaisResponses, domain.PartaiResponses{
			ID:            partai.ID,
			Name:          partai.Name,
			Chairman:      partai.Chairman,
			VisionMission: partai.VisionMission,
			Address:       partai.Address,
			Image:         partai.Image,
		})
	}
	
	return partaisResponses, nil
}

func (u *partaiUseCase) GetByID(id int64) (domain.PartaiResponse, error) {
	ctx := context.Background()
	partai, err := u.partaiRepository.GetByID(ctx, id)
	if err != nil {
		return domain.PartaiResponse{}, err
	}
	
	partaiResponse := domain.PartaiResponse{
		ID:            partai.ID,
		Name:          partai.Name,
		Chairman:      partai.Chairman,
		VisionMission: partai.VisionMission,
		Address:       partai.Address,
		Image:         partai.Image,
	}
	
	return partaiResponse, nil
}

func (u *partaiUseCase) Store(partai domain.Partai) (domain.Partai, error) {
	ctx := context.Background()
	partai, err := u.partaiRepository.Store(ctx, partai)
	if err != nil {
		return domain.Partai{}, err
	}
	
	return partai, nil
}

func (u *partaiUseCase) Update(id int64, partai domain.Partai) (domain.Partai, error) {
	ctx := context.Background()
	partai, err := u.partaiRepository.Update(ctx, id, partai)
	if err != nil {
		return domain.Partai{}, err
	}
	
	return partai, nil
}

func (u *partaiUseCase) Delete(id int64) error {
	ctx := context.Background()
	err := u.partaiRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	
	return nil
}
