package storage

import (
	"courses/internal/domain"
	"courses/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Courses interface {
	CreateCourse(course *domain.Course, userID int) (int, error)
	GetCourseByID(id int) (*domain.Course, error)
	UpdateCourse(id int, course *domain.UpdateCourse) error
	GetAllCourses() (*[]domain.Course, error)
	GetAllCoursesByTeacher(userID int) (*[]domain.Course, error)
}

type Lesson interface {
	CreateLesson(lesson *domain.Lesson) (int, error)
	GetLessonByName(name string) (*domain.Lesson, error)
	GetLessonByID(id int) (*domain.Lesson, error)
	UpdateLesson(id int, lessonData *domain.UpdateLesson) error
	UploadFile(lessonID int, fileName string, fileData []byte) error
	SendLessonForMarking(lessonID int) error
	GetAllDoneLesson() (*[]domain.Lesson, error)
	//TODO: check it
	GetAllDoneLessonByCourse(course int) (*[]domain.Lesson, error)
}

type Checklist interface {
	//TODO: add methods
}

type Storage struct {
	Lesson
	Courses
}

// TODO: init storage
func NewStorage(db *sqlx.DB, logg *slog.Logger) *Storage {
	return &Storage{
		Lesson:  postgres.NewLessonStorage(db, logg),
		Courses: postgres.NewCourseStorage(db, logg),
	}
}
