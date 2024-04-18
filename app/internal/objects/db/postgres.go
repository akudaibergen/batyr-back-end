package db

import (
	"app/internal/objects"
	"app/pkg/logging"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ objects.Storage = &db{}

type db struct {
	client *gorm.DB
	logger logging.Logger
}

func (d db) CancelPayments(uid uuid.UUID) (err error) {
	result := d.client.Table("users").Where("id = ?", uid).Update("is_paid", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d db) GetAllUsers() (data []objects.User, err error) {
	result := d.client.Table("users").Find(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (d db) SetUserPayment(uid uuid.UUID) (err error) {
	result := d.client.Table("users").Where("id = ?", uid).Update("is_paid", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d db) GetUserByUid(uid uuid.UUID) (data objects.User, err error) {
	result := d.client.Table("users").Where("id = ?", uid).First(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (d db) GetUsers() (data []objects.User, err error) {
	result := d.client.Table("users").Where("is_paid = true").Find(&data)
	if result.Error != nil {
		return data, result.Error
	}
	return data, nil
}

func (d db) CreateUser(dto objects.User) (data objects.User, err error) {
	result := d.client.Table("users").Create(&dto)
	if result.Error != nil {
		return data, result.Error
	}
	return dto, nil
}

func NewStorage(client *gorm.DB, logger logging.Logger) objects.Storage {
	return &db{
		client: client,
		logger: logger,
	}
}
