package database

import "github.com/gustavoarendt/jobtracker/internal/entities"

type UserDBInterface interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}

type CompanyDBInterface interface {
	Create(company *entities.Company) error
	FindById(id uint) (*entities.Company, error)
	FindByName(name string) (*entities.Company, error)
	FindAll() ([]entities.Company, error)
	Update(company *entities.Company) error
	Delete(company *entities.Company) error
}

type JobDBInterface interface {
	Create(job *entities.Job) error
	FindById(id uint) (*entities.Job, error)
	FindAll() ([]entities.Job, error)
	Update(job *entities.Job) error
	Delete(job *entities.Job) error
}
