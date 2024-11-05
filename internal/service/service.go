package service

import (
	"courses/internal/domain"
	"courses/internal/storage"
	"log/slog"
)

type Course interface {
	CreateCourse(course *domain.Course, userID int) (int, error)
	GetCourseByID(id int) (*domain.Course, error)
	UpdateCourse(id int, course *domain.UpdateCourse) error
	GetAllCourses() (*[]domain.Course, error)
	GetAllCoursesByTeacher(userID int) (*[]domain.Course, error)
	DeleteCourse(id int) error
	//AddStudentToCourse(courseID int, userID int) error
	//ApplyToCourse(courseID int, userID int) error
	//RespondToCourse(courseID int, userID int) error
}

type Lesson interface {
	CreateLesson(courseID int, lesson *domain.Lesson) (int, error)
	GetLessonByName(courseID int, name string) (*domain.Lesson, error)
	GetLessonByID(courseID int, id int) (*domain.Lesson, error)
	GetAllLessons(courseID int) (*[]domain.Lesson, error)
	UpdateLesson(courseID int, id int, lessonData *domain.UpdateLesson) error
	UploadFile(courseID int, lessonID int, fileName string, fileData []byte) error
	SendLessonForMarking(courseID int, lessonID int) error
	//GetAllLesson() (*[]domain.Lesson, error)
	//GetAllDoneLesson(status string) (*[]domain.Lesson, error)
	//GetAllDoneLessonByCourse(course int) (*[]domain.Lesson, error)
}

type Auth interface {
	ParseToken(token string) (int, error)
}

type Service struct {
	Lesson
	Course
	Auth
}

func NewService(storage *storage.Storage, logg *slog.Logger) *Service {
	return &Service{
		Lesson: NewLessonService(storage.Lesson),
		Course: NewCourseService(storage.Courses),
		Auth:   NewAuthService(logg),
	}
}
