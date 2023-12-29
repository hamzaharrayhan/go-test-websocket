package repository

import (
	"jti-test/internal/models"

	"gorm.io/gorm"
)

type ProviderRepository interface {
	Save(provider models.Provider) (models.Provider, error)
	FindAll() ([]models.Provider, error)
}

type providerRepository struct {
	db *gorm.DB
}

func NewProviderRepository(db *gorm.DB) *providerRepository {
	return &providerRepository{db}
}

func (p *providerRepository) Save(provider models.Provider) (models.Provider, error) {
	err := p.db.Save(&provider).Error

	if err != nil {
		return models.Provider{}, err
	}

	return provider, nil
}

func (p *providerRepository) FindAll() ([]models.Provider, error) {
	var providers []models.Provider

	err := p.db.Find(&providers).Error

	if err != nil {
		return nil, err
	}

	return providers, err
}
