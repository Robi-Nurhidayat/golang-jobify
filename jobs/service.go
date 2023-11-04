package jobs

type JobsServiceImpl interface {
	CreateJob(input JobsCreateInput)(Jobs,error)
	GetAllJobs()([]Jobs,error)
}

type JobsService struct {
	repository JobsRepositoryImpl
}

func NewJobsService(repository JobsRepositoryImpl) JobsServiceImpl {
	return &JobsService {
		repository: repository,
	}
}

func (s *JobsService) CreateJob(input JobsCreateInput)(Jobs,error) {
	jobs := Jobs{}

	jobs.Company = input.Company
	jobs.Position = input.Position
	jobs.Status =input.Status
	jobs.JobType = input.JobType
	jobs.JobLocation = input.JobLocation

	job,err := s.repository.CreateJob(jobs)

	if err != nil {
		return job,err
	}

	return job,nil
}

func (s *JobsService) GetAllJobs()([]Jobs,error) {
	jobs,err := s.repository.GetAllJobs()

	if err != nil {
		return jobs,err
	}

	return jobs,nil
}