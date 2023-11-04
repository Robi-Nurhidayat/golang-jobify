package jobs

import "gorm.io/gorm"

type JobsRepositoryImpl interface {
	CreateJob(jobs Jobs) (Jobs, error)
	GetAllJobs() ([]Jobs, error)
	UpdateJob(jobs Jobs) (Jobs, error)
	DeleteJob(id int) (Jobs, error)
}

type JobsRepository struct {
	db *gorm.DB
}

func NewJobsRepository(db *gorm.DB) JobsRepositoryImpl {
	return &JobsRepository{
		db: db,
	}
}

func (j *JobsRepository) CreateJob(jobs Jobs) (Jobs, error) {
	
	err := j.db.Create(&jobs).Error

	if err != nil {
		return jobs,err
	}

	return jobs,nil
}

func (j *JobsRepository) GetAllJobs() ([]Jobs, error) {
	
	var jobs []Jobs

	err := j.db.Find(&jobs).Error

	if err != nil {
		return jobs,err
	}

	return jobs,nil
}

func (j *JobsRepository) UpdateJob(jobs Jobs) (Jobs, error) {
	err := j.db.Save(&jobs).Error

	if err != nil {
		return jobs,err
	}

	return jobs,err
}

func (j *JobsRepository) DeleteJob(id int) (Jobs, error) {
	//TODO implement me
	panic("implement me")
}
