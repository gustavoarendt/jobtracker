package database

import (
	"github.com/gustavoarendt/jobtracker/internal/entities"
	"gorm.io/gorm"
)

type CompanyDB struct {
	DB *gorm.DB
}

func NewCompany(db *gorm.DB) *CompanyDB {
	return &CompanyDB{DB: db}
}

func (c *CompanyDB) Create(company *entities.Company) error {
	return c.DB.Create(company).Error
}

func (c *CompanyDB) FindById(id uint) (*entities.Company, error) {
	company := entities.Company{}
	if err := c.DB.Where("id = ?", id).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (c *CompanyDB) FindByName(name string) (*entities.Company, error) {
	company := entities.Company{}
	if err := c.DB.Where("name = ?", name).First(&company).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (c *CompanyDB) FindAll() ([]entities.Company, error) {
	var companies []entities.Company
	if err := c.DB.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (c *CompanyDB) Update(company *entities.Company) error {
	return c.DB.Save(company).Error
}

func (c *CompanyDB) Delete(company *entities.Company) error {
	return c.DB.Delete(company).Error
}
