package dto

import "time"

type User struct {
	UserName     string `json:"username"`
	UserLastName string `json:"lastname"`
	UserPassword string `json:"userpassword"`
}

type ResponsUser struct {
	Id           string    `json:"id"`
	UserName     string    `json:"username"`
	UserLastName string    `json:"lastname"`
	UserPassword string    `json:"userpassword"`
	Time         time.Time `json:"time"`
}

type LoginUser struct {
	UsersName    string `json:"username"`
	UserPassword string `json:"userpassword"`
}
