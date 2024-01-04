package jobs

type JobsFormatter struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Company     string `json:"company"`
	Position    string `json:"position"`
	Status      string `json:"status"`
	JobType     string `json:"job_type"`
	JobLocation string `json:"job_location"`
}

func FormatterJob(job Job) JobsFormatter {

	formatter := JobsFormatter{}

	formatter.Id = job.Id
	formatter.UserId = job.UserId
	formatter.Company = job.Company
	formatter.Position = job.Position
	formatter.Status = job.Status
	formatter.JobType = job.JobType
	formatter.JobLocation = job.JobLocation

	return formatter
}
