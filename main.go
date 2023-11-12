package main

import (
	"github.com/golang-jwt/jwt/v5"
	"jobify/auth"
	"jobify/handler"
	"jobify/helper"
	"jobify/jobs"
	"jobify/user"
	"log"
	"net/http"
	"strings"

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
	authService := auth.NewJwtService()

	// handler
	userHandler := handler.NewUserHandler(userService, authService)
	jobsHandler := handler.NewJobsHandler(jobService)

	r := gin.Default()
	api := r.Group("api/v1")
	//users
	api.POST("/user/register", userHandler.Register)
	api.POST("/user/login", userHandler.Login)
	api.GET("/user/get-all", userHandler.AllUsers)

	//jobs
	api.POST("/job", authMiddleware(authService, userService), jobsHandler.CreateJobs)
	api.GET("/jobs", authMiddleware(authService, userService), jobsHandler.GetAllJobs)
	api.GET("/jobs/user", authMiddleware(authService, userService), jobsHandler.GetAllJobsByUser)
	api.DELETE("/jobs/:id", authMiddleware(authService, userService), jobsHandler.DeleteJob)
	api.PUT("/jobs/:id", authMiddleware(authService, userService), jobsHandler.Update)

	r.Run()
}

func authMiddleware(authService auth.ServiceAuth, serviceUser user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {

			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""

		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			if err != nil {
				response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
		}

		userId := int(claim["user_id"].(float64))

		newUser, err := serviceUser.GetUserById(userId)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", newUser)
	}
}
