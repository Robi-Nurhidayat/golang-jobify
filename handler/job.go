package handler

import (
	"github.com/gin-gonic/gin"
	"jobify/helper"
	"jobify/jobs"
	"jobify/user"
	"net/http"
	"strconv"
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

		response := helper.ApiResponse("Please isi semua field", http.StatusFound, "failed", helper.FormatValidationError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userByAuth := c.MustGet("currentUser").(user.User)

	input.User = userByAuth
	job, err := h.service.CreateJob(input)
	if err != nil {

		response := helper.ApiResponse("Failed to create job", http.StatusFound, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Successfully create job", http.StatusOK, "success", jobs.FormatterJob(job))

	c.JSON(http.StatusOK, response)

}

func (h *JobsHandler) GetAllJobs(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	jobsAll, err := h.service.GetAllJobs(page, pageSize)
	if err != nil {
		response := helper.ApiResponse("Failed get All data", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
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

func (h *JobsHandler) GetAllJobsByUser(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)
	userId := int(currentUser.Id)
	jobsAll, err := h.service.GetAllJobsByUser(userId)
	if err != nil {
		response := helper.ApiResponse("Failed get All data", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
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
		response := helper.ApiResponse("Not found id", http.StatusFound, "failed", "tidak ada id")
		c.JSON(http.StatusBadRequest, response)

		return
	}

	job, err := h.service.GetById(input.Id)

	if err != nil {
		response := helper.ApiResponse("Not found id", http.StatusFound, "failed", "tidak ada id")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if job.Id == 0 {
		response := helper.ApiResponse("Not found id", http.StatusFound, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.DeleteJob(input)
	if err != nil {
		response := helper.ApiResponse("Not found id", http.StatusFound, "failed", "tidak ada id")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("successfully delete", http.StatusOK, "success", jobs.FormatterJob(job))
	c.JSON(http.StatusOK, response)

}

func (h *JobsHandler) Update(c *gin.Context) {

	var input jobs.JobId

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("Not found id", http.StatusFound, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)

		return
	}

	job, err := h.service.GetById(input.Id)

	if err != nil {
		response := helper.ApiResponse("Not found id", http.StatusFound, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if job.Id == 0 {
		response := helper.ApiResponse("Not found id", http.StatusFound, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	if currentUser.Id != job.UserId {
		response := helper.ApiResponse("Not matched id", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var jobInput jobs.JobsUpdateInput

	err = c.ShouldBindJSON(&jobInput)

	if err != nil {
		response := helper.ApiResponse("Not found id", http.StatusFound, "failed", helper.FormatValidationError(err))
		c.JSON(http.StatusBadRequest, response)

		return
	}

	job.Company = jobInput.Company
	job.Position = jobInput.Position
	job.Status = jobInput.Status
	job.JobType = jobInput.JobType
	job.JobLocation = jobInput.JobLocation

	jobUpdate, err := h.service.Update(job)
	if err != nil {
		response := helper.ApiResponse("failed update job", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("successfully update job", http.StatusOK, "success", jobs.FormatterJob(jobUpdate))
	c.JSON(http.StatusOK, response)

}
