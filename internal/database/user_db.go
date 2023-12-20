package database

import (
	"github.com/gustavoarendt/jobtracker/internal/entities"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *UserDB {
	return &UserDB{DB: DB}
}

func (u *UserDB) CreateUser(user *entities.User) error {
	return u.DB.Create(user).Error
}

func (u *UserDB) FindByEmail(email string) (*entities.User, error) {
	user := entities.User{}
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
