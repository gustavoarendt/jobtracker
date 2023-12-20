package entities

import "time"

type Company struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"type:varchar(50);not null" json:"name"`
	Description  string    `gorm:"type:varchar(2000)" json:"description"`
	Website_url  string    `gorm:"type:varchar(511)" json:"website_url"`
	Linkedin_url string    `gorm:"type:varchar(511)" json:"linkedin_url"`
	Image_url    string    `gorm:"type:varchar(511)" json:"image_url"`
	Created_at   time.Time `gorm:"type:timestamp" json:"created_at"`
}

func NewCompany(name, description, website_url, linkedin_url, image_url string) (*Company, error) {
	company := &Company{
		Name:         name,
		Description:  description,
		Website_url:  website_url,
		Linkedin_url: linkedin_url,
		Image_url:    image_url,
		Created_at:   time.Now(),
	}
	err := company.Validate()
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c *Company) Validate() error {
	if c.Name == "" {
		return ErrRequiredName
	}
	return nil
}
