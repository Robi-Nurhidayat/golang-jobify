package jobs

type JobsService interface {
	CreateJob(input JobsCreateInput) (Jobs, error)
	GetAllJobs() ([]Jobs, error)
	GetAllJobsByUser(userId int) ([]Jobs, error)
	DeleteJob(jobId JobId) (int, error)
	GetById(id int) (Jobs, error)
	Update(jobs Jobs) (Jobs, error)
}

type jobsService struct {
	repository JobsRepository
}

func NewJobsService(repository JobsRepository) *jobsService {
	return &jobsService{
		repository: repository,
	}
}

func (s *jobsService) CreateJob(input JobsCreateInput) (Jobs, error) {
	jobs := Jobs{}

	jobs.Company = input.Company
	jobs.Position = input.Position
	jobs.Status = input.Status
	jobs.JobType = input.JobType
	jobs.JobLocation = input.JobLocation
	jobs.UserId = input.User.Id

	job, err := s.repository.CreateJob(jobs)

	if err != nil {
		return job, err
	}

	return job, nil
}

func (s *jobsService) GetAllJobs() ([]Jobs, error) {
	jobs, err := s.repository.GetAllJobs()

	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (s *jobsService) GetAllJobsByUser(userId int) ([]Jobs, error) {
	jobs, err := s.repository.GetAllJobsByUser(userId)

	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (s *jobsService) DeleteJob(jobId JobId) (int, error) {
	id, err := s.repository.DeleteJob(jobId.Id)

	if err != nil {
		return id, err
	}

	return id, nil
}
func (s *jobsService) GetById(id int) (Jobs, error) {
	job, err := s.repository.GetById(id)

	if err != nil {
		return job, err
	}

	return job, nil
}

func (s *jobsService) Update(jobs Jobs) (Jobs, error) {
	job, err := s.repository.UpdateJob(jobs)

	if err != nil {
		return job, err
	}

	return job, nil
}
