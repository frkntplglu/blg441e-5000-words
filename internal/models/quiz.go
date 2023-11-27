package models

type Quiz struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Level string `json:"level"`
}
