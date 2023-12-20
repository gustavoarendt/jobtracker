package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompany_Instance(t *testing.T) {
	company, err := NewCompany("Test Company", "Test Description", "https://www.test.com", "https://www.linkedin.com/test", "https://www.test.com/image.png")
	assert.Nil(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, "Test Company", company.Name)
	assert.Equal(t, "Test Description", company.Description)
	assert.Equal(t, "https://www.test.com", company.Website_url)
	assert.Equal(t, "https://www.linkedin.com/test", company.Linkedin_url)
	assert.Equal(t, "https://www.test.com/image.png", company.Image_url)
	assert.NotNil(t, company.Created_at)
}

func TestCompany_ValidateNameRequired(t *testing.T) {
	company, err := NewCompany("", "Test Description", "https://www.test.com", "https://www.linkedin.com/test", "https://www.test.com/image.png")
	assert.NotNil(t, err)
	assert.Nil(t, company)
}
