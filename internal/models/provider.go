package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Provider struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Provider string    `gorm:"not null; unique"`
}

func (provider *Provider) BeforeCreate(tx *gorm.DB) (err error) {
	provider.ID = uuid.New()
	return
}
