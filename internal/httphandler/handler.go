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
	router.Route("/lessons", func(r chi.Router) {
		r.Post("/create", h.createLesson)
		r.Get("/name/{name}", h.getLessonByName)
		r.Get("/id/{id}", h.getLessonByID)
		r.Get("/done", h.getAllDoneLessons)
		r.Get("/done/{id}", h.getAllDoneLessons)
		r.Put("/update/{id}", h.updateLessonStatus)
		r.Post("/send/{id}", h.sendLessonForMarking)
	})

	return router
}
