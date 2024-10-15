package domain

import (
	"errors"
	"fmt"
)

type Lesson struct {
	LessonID     int     `json:"lesson_id" db:"lesson_id"`
	LessonType   string  `json:"lesson_type" db:"lesson_type"`
	Name         string  `json:"name" db:"lesson_name"`
	Description  string  `json:"description" db:"lesson_description"`
	LessonFile   *string `json:"lesson_file" db:"lesson_file_name"`
	FileContent  *[]byte `json:"file_content" db:"lesson_file_content"`
	LessonStatus string  `json:"lesson_status" db:"lesson_status"` //('Not start', 'send', 'rejected')
}

type UpdateLesson struct {
	LessonType  *string `json:"lesson_type" db:"lesson_type"` // ('lecture', 'practice')
	Name        *string `json:"name" db:"lesson_name"`
	Description *string `json:"description" db:"lesson_description"`
}

func (l Lesson) Validate() error {
	if l.Name == "" || l.Description == "" || l.LessonType == "" {
		return errors.New(fmt.Sprintf("invalid name: %s, description: %s, lesson_type: %s",
			l.Name, l.Description, l.LessonType))
	}
	return nil
}

func (l UpdateLesson) Validate() error {
	if l.LessonType == nil && l.Name == nil && l.Description == nil {
		return errors.New(fmt.Sprintf("invalid lesson id: %d or status: %d or lesson_desc: %d",
			l.LessonType, l.Name, l.Description))
	}
	return nil
}

type TeacherChecklist struct {
	ID        int    `json:"id" db:"id"`
	TeacherID int    `json:"teacher_id" db:"teacher_id"`
	LessonID  int    `json:"lesson_id" db:"lesson_id"`
	Status    string `json:"status" db:"status"` //('On checking', 'done', 'rejected')
	Homework  string `json:"homework" db:"homework"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

//{
//	"name": "check",
//	"description": "check description"
//	"status": "done"
//}
