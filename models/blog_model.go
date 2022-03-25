package models

type Blog struct {
	Title       string `json:"title"`
	Datetime    string `json:"datetime"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
