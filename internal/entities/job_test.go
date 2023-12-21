package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJob_Instance(t *testing.T) {
	job, err := NewJob("Test Job", "Test Description", "Refused", "BRL", "Portuguese", 1, 1, 1000.00, 1)
	assert.Nil(t, err)
	assert.NotNil(t, job)
	assert.Equal(t, "Test Job", job.Name)
	assert.Equal(t, "Test Description", job.Description)
	assert.Equal(t, "Refused", job.Status)
	assert.Equal(t, "BRL", job.Currency)
	assert.Equal(t, "Portuguese", job.Language)
	assert.Equal(t, uint64(1), job.Id_company)
	assert.Equal(t, 1000.00, job.Expected_salary)
	assert.Equal(t, 1, job.Interest)
	assert.NotNil(t, job.Created_at)
	assert.NotNil(t, job.Updated_at)
}

func TestJob_ValidateNameRequired(t *testing.T) {
	job, err := NewJob("", "Test Description", "Refused", "BRL", "Portuguese", 1, 1, 1000.00, 1)
	assert.NotNil(t, err)
	assert.Nil(t, job)
}

func TestJob_ValidateCurrencyLength(t *testing.T) {
	job, err := NewJob("Test Job", "Test Description", "Refused", "Real Brasileiro", "Portuguese", 1, 1, 1000.00, 1)
	assert.NotNil(t, err)
	assert.Nil(t, job)
}

func TestJob_ValidateInterestRange(t *testing.T) {
	job, err := NewJob("Test Job", "Test Description", "Refused", "Real Brasileiro", "Portuguese", 1, 1, 1000.00, 10)
	assert.NotNil(t, err)
	assert.Nil(t, job)
}

func TestJob_ValidateCompanyIDRequired(t *testing.T) {
	job, err := NewJob("Test Job", "Test Description", "Refused", "Real Brasileiro", "Portuguese", 1, 1, 1000.00, 10)
	assert.NotNil(t, err)
	assert.Nil(t, job)
}

func TestJob_ValidateStatusRequired(t *testing.T) {
	job, err := NewJob("Test Job", "Test Description", "", "BRL", "Portuguese", 1, 1, 1000.00, 1)
	assert.NotNil(t, err)
	assert.Nil(t, job)
}
