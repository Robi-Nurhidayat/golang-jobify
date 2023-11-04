package jobs

type JobsFormatter struct {
	Id          int    `json:"id"`
	Company     string `json:"company"`
	Position    string `json:"position"`
	Status      string `json:"status"`
	JobType     string `json:"job_type"`
	JobLocation string `json:"job_location"`
}

func FormatterJob(job Jobs) JobsFormatter {

	formatter := JobsFormatter{}

	formatter.Id = job.Id
	formatter.Company = job.Company
	formatter.Position = job.Position
	formatter.Status = job.Status
	formatter.JobType = job.JobType
	formatter.JobLocation = job.JobLocation

	return formatter
}