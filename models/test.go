package models

type ChoiceQuestion struct {
	QuestionID int64  `json:"question_id" db:"question_id"  `
	Content    string `json:"content" db:"content"  binding:"required"`
	Options    string `json:"options" db:"options"  binding:"required"`
	Type       string `json:"type" db:"type"  binding:"required"`
	Answer     string `json:"answer" db:"answer"  binding:"required"`
}

type GapFilling struct {
	QuestionID int64  `json:"question_id" db:"question_id" `
	Content    string `json:"content" db:"content"  binding:"required"`
	Answer     string `json:"answer" db:"answer"  binding:"required"`
	Type       string `json:"type" db:"type"  binding:"required"`
}

type Test struct {
	TestID         int64            `json:"test_id"`
	ChoiceQuestion []ChoiceQuestion `json:"choice_question"`
	GapFilling     []GapFilling     `json:"gap_filling"`
}

type Answers struct {
	QuestionID int64  `json:"question_id"`
	Answer     string `json:"answer"`
}

type ResponseAnswers struct {
	QuestionID int64  `json:"question_id"`
	Answer     string `json:"answer"`
	Result     bool
}

type ChoiceQuestionWithoutAnswer struct {
	QuestionID int64  `json:"question_id" db:"question_id"`
	Content    string `json:"content" db:"content"`
	Options    string `json:"options" db:"options"`
}

type GapFillingWithoutAnswer struct {
	QuestionID int64  `json:"question_id" db:"question_id"`
	Content    string `json:"content" db:"content"`
}

