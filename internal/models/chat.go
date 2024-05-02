package models

type Chat struct {
	ID         int `json:"id" db:"id"`
	FirstUser  int `json:"first_user" db:"first_user"`
	SecondUser int `json:"second_user" db:"second_user"`
}
