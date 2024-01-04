package user

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Avatar    string
	Password  string
	LastName  string
	Location  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}


func (User) TableName() string {
	return "users"
}