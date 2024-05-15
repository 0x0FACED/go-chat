package models

type Chat struct {
	ID         int `json:"id" db:"id"`
	FirstUser  int `json:"first_user" db:"first_user_id"`
	SecondUser int `json:"second_user" db:"second_user_id"`
}
