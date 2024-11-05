package httphandler

import (
	"courses/internal/httphandler/mw/logger"
	"courses/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log/slog"
)

type Handler struct {
	service *service.Service
	logg    *slog.Logger
}

func NewHandler(service *service.Service, logg *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logg:    logg,
	}
}

func (h *Handler) InitRoutes(logg *slog.Logger) chi.Router {
	router := chi.NewRouter()

	router.Use(logger.New(logg))
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.RealIP)

	router.Route("/courses", func(r chi.Router) {
		r.Use(h.userIdentity)
		r.Post("/create", h.createCourse)
		r.Get("/id/{id}", h.getCourseByID)
		r.Put("/update/{id}", h.updateCourse)
		r.Get("/all", h.getAllCourses)
		r.Get("/all/{id}", h.getAllCoursesByTeacher)
		r.Delete("/delete/{id}", h.deleteCourse)
		//r.Post("/apply", h.applyToCourse)
		//r.Post("/add", h.addStudentToCourse)
		r.Route("/{course_id}/lessons", func(insideRouter chi.Router) {
			insideRouter.Post("/create", h.createLesson)
			insideRouter.Get("/name/{name}", h.getLessonByName)
			insideRouter.Get("/id/{id}", h.getLessonByID)
			insideRouter.Get("/all", h.getAllLessons)
			insideRouter.Put("/update/{id}", h.updateLesson)
			insideRouter.Post("/upload/{id}/{filename}", h.uploadFile)
			insideRouter.Post("/send/{id}", h.sendLessonForMarking)
			//TODO: check later
			//r.Get("/done", h.getAllDoneLessons)
			//TODO: fix it
			//r.Get("/done/course/{id}", h.getAllDoneLessonsByCourse)
		})
	})

	return router
}
