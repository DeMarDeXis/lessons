package service

import (
	"courses/internal/domain"
	"courses/internal/storage"
)

type Course interface {
	CreateCourse(course *domain.Course, userID int) (int, error)
	GetCourseByID(id int) (*domain.Course, error)
	UpdateCourse(id int, course *domain.UpdateCourse) error
	GetAllCourses() (*[]domain.Course, error)
	GetAllCoursesByTeacher(userID int) (*[]domain.Course, error)
	//AddStudentToCourse(courseID int, userID int) error
	//ApplyToCourse(courseID int, userID int) error
	//RespondToCourse(courseID int, userID int) error
}

type Lesson interface {
	CreateLesson(lesson *domain.Lesson) (int, error)
	GetLessonByName(name string) (*domain.Lesson, error)
	GetLessonByID(id int) (*domain.Lesson, error)
	UpdateLesson(id int, lessonData *domain.UpdateLesson) error
	UploadFile(lessonID int, fileName string, fileData []byte) error
	SendLessonForMarking(lessonID int) error
	GetAllDoneLesson() (*[]domain.Lesson, error)
	GetAllDoneLessonByCourse(course int) (*[]domain.Lesson, error)
}

type Service struct {
	Lesson
	Course
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Lesson: NewLessonService(storage.Lesson),
		Course: NewCourseService(storage.Courses),
	}
}
