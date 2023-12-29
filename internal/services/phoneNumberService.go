package services

import (
	"fmt"
	"jti-test/internal/helpers"
	"jti-test/internal/models"
	"jti-test/internal/repository"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type PhoneNumberService interface {
	CreatePhoneNumber(input helpers.PhoneNumberInput) (models.PhoneNumber, error)
	Delete(ID string) (models.PhoneNumber, error)
	Update(ID string, input helpers.PhoneNumberInput) (models.PhoneNumber, error)
	GetAll() ([]models.PhoneNumber, error)
}

type phoneNumberService struct {
	phoneNumberRepository repository.PhoneNumberRepository
}

func NewPhoneNumberService(phoneNumberRepository repository.PhoneNumberRepository) *phoneNumberService {
	return &phoneNumberService{phoneNumberRepository}
}

func (p *phoneNumberService) CreatePhoneNumber(input helpers.PhoneNumberInput) (models.PhoneNumber, error) {
	_, err := strconv.Atoi(input.Number)
	if err != nil {
		return models.PhoneNumber{}, err
	}

	newPhoneNumber := models.PhoneNumber{
		Number:     input.Number,
		ProviderID: uuid.MustParse(input.ProviderID),
	}

	createNewPhoneNumber, err := p.phoneNumberRepository.Save(newPhoneNumber)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return models.PhoneNumber{}, err
	}

	return createNewPhoneNumber, nil
}

func (p *phoneNumberService) Delete(ID string) (models.PhoneNumber, error) {
	phoneNumber, err := p.phoneNumberRepository.FindByID(ID)

	if err != nil {
		return models.PhoneNumber{}, err
	}

	if phoneNumber.ID == uuid.Nil {
		return models.PhoneNumber{}, nil
	}

	Deleted, err := p.phoneNumberRepository.Delete(ID)

	if err != nil {
		return models.PhoneNumber{}, err
	}

	return Deleted, nil
}

func (p *phoneNumberService) Update(ID string, input helpers.PhoneNumberInput) (models.PhoneNumber, error) {
	_, err := strconv.Atoi(input.Number)
	if err != nil {
		return models.PhoneNumber{}, err
	}

	Result, err := p.phoneNumberRepository.FindByID(ID)

	if err != nil {
		return models.PhoneNumber{}, err
	}

	if Result.ID == uuid.Nil {
		return models.PhoneNumber{}, nil
	}

	updated := models.PhoneNumber{
		Number:     string(helpers.Encrypt([]byte(fmt.Sprint(input.Number)), []byte(viper.GetString("key")))),
		ProviderID: uuid.MustParse(fmt.Sprint(input.ProviderID)),
	}

	phoneNumberUpdate, err := p.phoneNumberRepository.Update(updated, ID)

	if err != nil {
		return models.PhoneNumber{}, err
	}

	return phoneNumberUpdate, nil
}

func (s *phoneNumberService) GetAll() ([]models.PhoneNumber, error) {
	res, err := s.phoneNumberRepository.FindAll()

	if err != nil {
		return []models.PhoneNumber{}, err
	}

	newRes := []models.PhoneNumber{}
	for _, v := range res {
		buf := models.PhoneNumber{
			ID:         v.ID,
			Number:     string(helpers.Decrypt(v.Number, []byte(viper.GetString("key")))),
			ProviderID: v.ProviderID,
		}
		newRes = append(newRes, buf)
	}

	return newRes, nil
}
