package database

import (
	"github.com/gustavoarendt/jobtracker/internal/entities"
	"gorm.io/gorm"
)

type JobDB struct {
	DB *gorm.DB
}

func NewJob(db *gorm.DB) *JobDB {
	return &JobDB{DB: db}
}

func (j *JobDB) Create(job *entities.Job) error {
	return j.DB.Create(job).Error
}

func (j *JobDB) FindById(id uint64) (*entities.Job, error) {
	job := entities.Job{}
	if err := j.DB.Where("id = ?", id).First(&job).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (j *JobDB) FindAll() ([]entities.Job, error) {
	var jobs []entities.Job
	if err := j.DB.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (j *JobDB) Update(job *entities.Job) error {
	return j.DB.Save(job).Error
}

func (j *JobDB) Delete(id string) error {
	return j.DB.Where("id = ?", id).Delete(&entities.Job{}).Error
}
