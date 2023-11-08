package main

import (
	"jobify/handler"
	"jobify/jobs"
	"jobify/user"
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
	userRepository := user.NewUserRepository(db)
	jobRepository := jobs.NewJobsRepository(db)

	// service
	userService := user.NewUserService(userRepository)
	jobService := jobs.NewJobsService(jobRepository)

	// handler
	userHandler := handler.NewUserHandler(userService)
	jobsHandler := handler.NewJobsHandler(jobService)

	r := gin.Default()
	api := r.Group("api/v1")
	//users
	api.POST("/user/register", userHandler.Register)

	//jobs
	api.POST("/job", jobsHandler.CreateJobs)
	api.GET("/jobs", jobsHandler.GetAllJobs)
	api.DELETE("/jobs/:id", jobsHandler.DeleteJob)
	api.PUT("/jobs/:id", jobsHandler.Update)

	r.Run()
}
