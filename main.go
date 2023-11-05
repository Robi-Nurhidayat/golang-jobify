package main

import (
	"jobify/handler"
	"jobify/jobs"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(localhost:3306)/jobify?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// repository
	// jobsRepository := jobs.NewJobsRepository(db)
	jobRepository := jobs.NewJobsRepository(db)

	// service
	// jobsService := jobs.NewJobsService(jobsRepository)

	jobService := jobs.NewJobsService(jobRepository)

	// handler

	// jobsHandler := handler.NewJobsHandler(jobsService)

	// jobsHandler := handler.NewJobsHandler(jobsService)
	jobsHandler := handler.NewJobsHandler(jobService)

	r := gin.Default()
	api := r.Group("api/v1")
	api.POST("/job", jobsHandler.CreateJobs)
	api.GET("/jobs", jobsHandler.GetAllJobs)
	api.DELETE("/jobs/:id", jobsHandler.DeleteJob)
	r.Run()
}
