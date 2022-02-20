package models

type Contact struct {
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	Text      string `json:"text"`
}
