package models

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	PassHash string `json:"pass_hash"`
}