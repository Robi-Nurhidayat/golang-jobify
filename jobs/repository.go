package jobs

import "gorm.io/gorm"

type JobsRepository interface {
	CreateJob(jobs Job) (Job, error)
	GetAllJobs(page, pageSize int) ([]Job, error)
	GetAllJobsByUser(userId int) ([]Job, error)
	UpdateJob(jobs Job) (Job, error)
	DeleteJob(id int) (int, error)
	GetById(id int) (Job, error)
}

type jobsRepository struct {
	db *gorm.DB
}

func NewJobsRepository(db *gorm.DB) JobsRepository {
	return &jobsRepository{
		db: db,
	}
}

func (j *jobsRepository) CreateJob(jobs Job) (Job, error) {

	err := j.db.Create(&jobs).Error

	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (j *jobsRepository) GetAllJobs(page, pageSize int) ([]Job, error) {

	var jobs []Job

	offset := (page - 1) * pageSize

	//err := j.db.Find(&jobs).Error

	err := j.db.Limit(pageSize).Offset(offset).Find(&jobs).Error

	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (j *jobsRepository) GetAllJobsByUser(userId int) ([]Job, error) {

	var jobs []Job

	err := j.db.Where("user_id = ?", userId).Find(&jobs).Error

	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (j *jobsRepository) UpdateJob(jobs Job) (Job, error) {
	err := j.db.Save(&jobs).Error

	if err != nil {
		return jobs, err
	}

	return jobs, err
}

func (j *jobsRepository) DeleteJob(id int) (int, error) {
	var jobs Job
	err := j.db.Where("id = ?", id).Delete(&jobs).Error

	if err != nil {
		return id, err
	}

	return id, nil
}
func (j *jobsRepository) GetById(id int) (Job, error) {
	var job Job
	err := j.db.Where("id = ?", id).Find(&job).Error

	if err != nil {
		return job, err
	}

	return job, nil
}
