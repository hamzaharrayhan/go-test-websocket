package repository

import (
	"fmt"
	"jti-test/internal/models"

	"gorm.io/gorm"
)

type PhoneNumberRepository interface {
	Save(phoneNumber models.PhoneNumber) (models.PhoneNumber, error)
	Delete(ID string) (models.PhoneNumber, error)
	Update(phoneNumber models.PhoneNumber, ID string) (models.PhoneNumber, error)
	FindAll() ([]models.PhoneNumber, error)
	FindByID(ID string) (models.PhoneNumber, error)
}

type phoneNumberRepository struct {
	db *gorm.DB
}

func NewPhoneNumberRepository(db *gorm.DB) *phoneNumberRepository {
	return &phoneNumberRepository{db}
}

func (p *phoneNumberRepository) Save(phoneNumber models.PhoneNumber) (models.PhoneNumber, error) {
	err := p.db.Save(&phoneNumber).Error

	if err != nil {
		fmt.Printf("%+v\n", err)
		return models.PhoneNumber{}, err
	}

	return phoneNumber, nil
}

func (p *phoneNumberRepository) Update(phoneNumber models.PhoneNumber, ID string) (models.PhoneNumber, error) {
	err := p.db.Where("id = ?", ID).Updates(&phoneNumber).Error

	if err != nil {
		return models.PhoneNumber{}, err
	}

	return phoneNumber, nil
}

func (p *phoneNumberRepository) Delete(ID string) (models.PhoneNumber, error) {
	err := p.db.Where("id = ?", ID).Delete(&models.PhoneNumber{}).Error

	if err != nil {
		return models.PhoneNumber{}, err
	}

	return models.PhoneNumber{}, nil
}

func (p *phoneNumberRepository) FindAll() ([]models.PhoneNumber, error) {
	var phoneNumbers []models.PhoneNumber

	err := p.db.Find(&phoneNumbers).Error

	if err != nil {
		return nil, err
	}

	return phoneNumbers, err
}

func (p *phoneNumberRepository) FindByID(ID string) (models.PhoneNumber, error) {
	var phoneNumber models.PhoneNumber
	err := p.db.First(&phoneNumber, "id = ?", ID).Error

	if err != nil {
		return models.PhoneNumber{}, err
	}

	return phoneNumber, nil
}
