package models

type SavedWord struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	WordId int `json:"wordId"`
}
