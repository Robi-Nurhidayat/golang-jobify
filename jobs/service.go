package jobs

type JobsService interface {
	CreateJob(input JobsCreateInput) (Job, error)
	GetAllJobs(page, pageSize int) ([]Job, error)
	GetAllJobsByUser(userId int) ([]Job, error)
	DeleteJob(jobId JobId) (int, error)
	GetById(id int) (Job, error)
	Update(jobs Job) (Job, error)
}

type jobsService struct {
	repository JobsRepository
}

func NewJobsService(repository JobsRepository) *jobsService {
	return &jobsService{
		repository: repository,
	}
}

func (s *jobsService) CreateJob(input JobsCreateInput) (Job, error) {
	jobs := Job{}

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

func (s *jobsService) GetAllJobs(page, pageSize int) ([]Job, error) {
	jobs, err := s.repository.GetAllJobs(page, pageSize)

	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (s *jobsService) GetAllJobsByUser(userId int) ([]Job, error) {
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
func (s *jobsService) GetById(id int) (Job, error) {
	job, err := s.repository.GetById(id)

	if err != nil {
		return job, err
	}

	return job, nil
}

func (s *jobsService) Update(jobs Job) (Job, error) {
	job, err := s.repository.UpdateJob(jobs)

	if err != nil {
		return job, err
	}

	return job, nil
}
