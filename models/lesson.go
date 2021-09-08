package models

import "time"

type Lesson struct {
	LessonId  int64     `json:"lesson_id" db:"lesson_id"`
	ClassId   int64     `json:"class_id" db:"class_id"`
	TeacherId int64     `json:"teacher_id" db:"teacher_id"`
	Homework  string    `json:"homework" db:"homework"`
	StartTime time.Time `json:"-" db:"start_time"`
	EndTime   time.Time `json:"-" db:"end_time"`
}
