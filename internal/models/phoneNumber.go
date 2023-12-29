package models

import (
	"jti-test/internal/helpers"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type PhoneNumber struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	Number     string    `gorm:"not null;uniqueIndex:unique_code_idx"`
	ProviderID uuid.UUID
	Provider   Provider
}

func (phoneNumber *PhoneNumber) BeforeCreate(tx *gorm.DB) (err error) {
	phoneNumber.ID = uuid.New()

	key := viper.GetString("key")
	number := helpers.Encrypt([]byte(phoneNumber.Number), []byte(key))
	phoneNumber.Number = string(number)

	return
}
