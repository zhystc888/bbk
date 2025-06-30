package model

type AdminBase struct {
	UserID int    `json:"userId" orm:"user_id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
