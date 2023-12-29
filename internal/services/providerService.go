package services

import (
	"jti-test/internal/models"
	"jti-test/internal/repository"
)

type ProviderService interface {
	GetAll() ([]models.Provider, error)
}

type providerService struct {
	providerRepository repository.ProviderRepository
}

func NewProviderService(providerRepository repository.ProviderRepository) *providerService {
	return &providerService{providerRepository}
}

func (s *providerService) GetAll() ([]models.Provider, error) {
	res, err := s.providerRepository.FindAll()

	if err != nil {
		return []models.Provider{}, err
	}

	return res, nil
}
