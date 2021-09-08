package models

import "time"

type LessonInfo struct {
	LessonId     int64     `json:"lesson_id" db:"lesson_id"`
	StudentId    int64     `json:"student_id" db:"student_id"`
	StudentName  string    `json:"student_name" db:"student_name"`
	SubmitStatus int       `json:"submit_status" db:"submit_status"`
	SubmitTime   time.Time `json:"submit_time" db:"submit_time"`
	SignInStatus int       `json:"signin_status" db:"signin_status"`
	SignInTime   time.Time `json:"signin_time" db:"signin_time"`
}
