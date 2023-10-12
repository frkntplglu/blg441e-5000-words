package models

type Word struct {
	Id          int    `json:"id"`
	Vocabulary  string `json:"vocabulary"`
	Definition  string `json:"definition"`
	Sentence    string `json:"sentence"`
	Translation string `json:"translation"`
	Level       string `json:"level"`
}

type Level string

const (
	Elementary   Level = "A1-A2"
	Intermediate Level = "B1-B2"
	Advance      Level = "C1-C2"
)
