package objects

import "github.com/google/uuid"

type Storage interface {
	CreateUser(dto User) (data User, err error)
	GetUsers() (data []User, err error)
	GetAllUsers() (data []User, err error)
	GetUserByUid(uid uuid.UUID) (data User, err error)
	SetUserPayment(uid uuid.UUID) (err error)
	CancelPayments(uid uuid.UUID) (err error)
}
