package jobs

type JobsCreateInput struct {
	
	Id          int    `json:"id"`
	Company     string `json:"company"`
	Position    string `json:"position"`
	Status      string `json:"status"`
	JobType     string `json:"job_type"`
	JobLocation string `json:"job_location"`
}