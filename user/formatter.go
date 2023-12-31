package user

type UserFormatter struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	LastName string `json:"last_name"`
	Location string `json:"location"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func FormatterUser(user User, token string) UserFormatter {
	formatter := UserFormatter{}

	formatter.Id = user.Id
	formatter.Name = user.Name
	formatter.Email = user.Email
	formatter.Avatar = user.Avatar
	formatter.LastName = user.LastName
	formatter.Location = user.Location
	formatter.Role = user.Role
	formatter.Token = token

	return formatter
}
