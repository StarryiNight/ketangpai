package models

type Class struct {
	ClassID   int64 `json:"class_id" db:"class_id"`
	ClassName string `json:"class_name" db:"class_name"`
}