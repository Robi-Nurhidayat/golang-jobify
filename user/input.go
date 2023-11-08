package user

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"email" binding:"required"`
}

type RegisterInput struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Location string `json:"location" binding:"required"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
}
