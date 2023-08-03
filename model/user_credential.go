package model

type UserCredential struct {
	Id       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password,omitempty"`
	IsActive bool   `json:"isActive"`
}
