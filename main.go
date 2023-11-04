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
	jobsRepository := jobs.NewJobsRepository(db)

	// service
	jobsService := jobs.NewJobsService(jobsRepository)

	// handler

	// jobsHandler := handler.NewJobsHandler(jobsService)

	jobsHandler := handler.NewJobsHandler(jobsService)

	r := gin.Default()
	api := r.Group("api/v1")
	api.POST("/job", jobsHandler.CreateJobs)
	api.GET("/jobs", jobsHandler.GetAllJobs)
	r.Run()
}
