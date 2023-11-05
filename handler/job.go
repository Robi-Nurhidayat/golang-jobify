package handler

import (
	"jobify/helper"
	"jobify/jobs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobsHandler struct {
	service jobs.JobsService
}

func NewJobsHandler(service jobs.JobsService) *JobsHandler {
	return &JobsHandler{
		service: service,
	}
}

func (h *JobsHandler) CreateJobs(c *gin.Context) {

	var input jobs.JobsCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		c.JSON(http.StatusBadRequest, nil)
		return
	}

	job, err := h.service.CreateJob(input)
	if err != nil {

		c.JSON(http.StatusBadRequest, nil)
		return
	}

	response := helper.ApiResponse("Successfully create job", http.StatusOK, "success", jobs.FormatterJob(job))

	c.JSON(http.StatusOK, response)

}

func (h *JobsHandler) GetAllJobs(c *gin.Context) {

	jobsAll, err := h.service.GetAllJobs()
	if err != nil {

		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var jobFormat []jobs.JobsFormatter

	for _, job := range jobsAll {
		temp := jobs.FormatterJob(job)
		jobFormat = append(jobFormat, temp)
	}
	response := helper.ApiResponse("Successfully create job", http.StatusOK, "success", jobFormat)

	c.JSON(http.StatusOK, response)

}

func (h *JobsHandler) DeleteJob(c *gin.Context) {

	var input jobs.JobId
	err := c.ShouldBindUri(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if input.Id == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	idByJob, err := h.service.DeleteJob(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	job,err := h.service.GetById(idByJob)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	response := helper.ApiResponse("successfully delete", http.StatusOK, "success", jobs.FormatterJob(job))
	c.JSON(http.StatusOK, response)

}
