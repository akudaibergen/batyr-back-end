package objects

import (
	"app/pkg/logging"
	"github.com/google/uuid"
)

var _ Service = &service{}

type service struct {
	storage Storage
	logger  logging.Logger
}

func (s service) CancelPayments(uid uuid.UUID) (err error) {
	return s.storage.CancelPayments(uid)
}

func (s service) GetAllUsers() (data []User, err error) {
	return s.storage.GetAllUsers()
}

func (s service) SetUserPayment(uid uuid.UUID) (err error) {
	return s.storage.SetUserPayment(uid)
}

func (s service) GetUserByUid(uid uuid.UUID) (data User, err error) {
	return s.storage.GetUserByUid(uid)
}

func (s service) GetUsers() (result []User, err error) {
	return s.storage.GetUsers()
}

func (s service) CreateUser(dto User) (data User, err error) {
	return s.storage.CreateUser(dto)
}

func NewService(storage Storage, logger logging.Logger) Service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

type Service interface {
	CreateUser(dto User) (data User, err error)
	GetUserByUid(uid uuid.UUID) (data User, err error)
	GetUsers() (data []User, err error)
	GetAllUsers() (data []User, err error)
	SetUserPayment(uid uuid.UUID) (err error)
	CancelPayments(uid uuid.UUID) (err error)
}
