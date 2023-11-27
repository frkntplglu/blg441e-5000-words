package models

type Question struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	OptionA string `json:"option_a"`
	OptionB string `json:"option_b"`
	OptionC string `json:"option_c"`
	OptionD string `json:"option_d"`
	Answer  string `json:"answer"`
	QuizId  int    `json:"quiz_id"`
}

type Answer struct {
	Answer string `json:"answer"`
}
