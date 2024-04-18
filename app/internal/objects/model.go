package objects

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	FirstName     string    `gorm:"column:firstname" json:"firstname"`
	Surname       string    `gorm:"column:surname" json:"surname"`
	Age           string    `gorm:"column:age" json:"age"`
	Region        string    `gorm:"column:region" json:"region"`
	Email         string    `gorm:"column:email" json:"email"`
	PhoneNumber   string    `gorm:"column:phone_number" json:"phone_number"`
	Schedule      string    `gorm:"column:schedule" json:"schedule"`
	Experience    string    `gorm:"column:experience" json:"experience"`
	AboutYourself string    `gorm:"column:about_yourself" json:"about_yourself"`
	DesiredSalary string    `gorm:"column:desired_salary" json:"desired_salary"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}
