package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	Email      string    `gorm:"type:varchar(255);not null" json:"email"`
	Password   string    `gorm:"type:varchar(511);not null" json:"-"`
	Created_at time.Time `gorm:"type:timestamp" json:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:       name,
		Email:      email,
		Password:   string(hash),
		Created_at: time.Now(),
	}, nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
