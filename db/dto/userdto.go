package dto

import "time"

type User struct{
	UserName string `json:"username"`
	UserLastName  string `json:"userlastname"`
}

type ResponsUser struct{
	Id string `json:"id"`
	UserName string `json:"username"`
	UserLastName  string `json:"userlastname"`
	Time time.Time `json:"time"`
}