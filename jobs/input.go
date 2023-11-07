package jobs

type JobsCreateInput struct {
	Id          int    `json:"id"`
	Company     string `json:"company" binding:"required"`
	Position    string `json:"position" binding:"required"`
	Status      string `json:"status" binding:"required"`
	JobType     string `json:"job_type" binding:"required"`
	JobLocation string `json:"job_location" binding:"required"`
}

type JobsUpdateInput struct {
	Id          int    `json:"id"`
	Company     string `json:"company" binding:"required"`
	Position    string `json:"position" binding:"required"`
	Status      string `json:"status" binding:"required"`
	JobType     string `json:"job_type" binding:"required"`
	JobLocation string `json:"job_location" binding:"required"`
}

type JobId struct {
	Id int `uri:"id"`
}
