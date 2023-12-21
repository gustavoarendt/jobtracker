package entities

import (
	"errors"
	"time"
)

var (
	ErrCurrencyLength    = errors.New("length of currency must be 3 characters")
	ErrInterestRange     = errors.New("interest must be between 1 and 5")
	ErrRequiredName      = errors.New("name is required")
	ErrRequiredCompanyID = errors.New("company id is required")
	ErrRequiredStatus    = errors.New("status is required")
)

type Job struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string    `gorm:"type:varchar(50);not null" json:"name"`
	Description     string    `gorm:"type:varchar(3000)" json:"description"`
	Status          string    `gorm:"type:varchar(50);not null" json:"status"`
	Currency        string    `gorm:"type:varchar(3)" json:"currency"`
	Language        string    `gorm:"type:varchar(50)" json:"language"`
	Id_company      uint64    `gorm:"type:numeric(20);not null" json:"id_company"`
	Id_user         uint64    `gorm:"type:numeric(20);not null" json:"id_user"`
	Expected_salary float64   `gorm:"type:numeric(15,2)" json:"expected_salary"`
	Interest        int       `gorm:"type:numeric(1);not null" json:"interest"`
	Created_at      time.Time `gorm:"type:timestamp" json:"created_at"`
	Updated_at      time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func NewJob(name, description, status, currency, language string, id_company, id_user uint64, expected_salary float64, interest int) (*Job, error) {
	job := &Job{
		Name:            name,
		Description:     description,
		Expected_salary: expected_salary,
		Interest:        interest,
		Status:          status,
		Currency:        currency,
		Language:        language,
		Id_company:      id_company,
		Id_user:         id_user,
		Created_at:      time.Now(),
		Updated_at:      time.Now(),
	}
	err := job.Validate()
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (j *Job) Validate() error {
	if len(j.Currency) > 3 {
		return ErrCurrencyLength
	}
	if j.Interest < 1 || j.Interest > 5 {
		return ErrInterestRange
	}
	if j.Name == "" {
		return ErrRequiredName
	}
	if j.Id_company <= 0 {
		return ErrRequiredCompanyID
	}
	if j.Status == "" {
		return ErrRequiredStatus
	}
	return nil
}
