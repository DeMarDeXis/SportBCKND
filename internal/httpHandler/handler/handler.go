package handler

import (
	"github.com/DeMarDeXis/VProj/internal/httpHandler/handler/mw/logger"
	"github.com/DeMarDeXis/VProj/internal/httpHandler/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
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

	//router.Use(middleware.Logger)
	router.Use(logger.New(logg))
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.RealIP)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	router.Route("/group", func(r chi.Router) {
		r.Use(h.userIdentity)
		r.Post("/lists", h.createList)
		//	r.Get("/lists", h.getAllLists)
		//	r.Get("/lists/{id}", h.getListByID)
		//	r.Put("/lists/{id}", h.updateList)
		//	r.Delete("/lists/{id}", h.deleteList)
	})
	return router
}

//router.Route("/url", func(r chi.Router) {
//	r.Use(middleware.BasicAuth("url-shortener", map[string]string{
//		cfg.HTTPServer.User: cfg.HTTPServer.Password,
//	}))
//
//	r.Post("/", save.New(log, storage))
//
//	r.Delete("/{alias}", delete.New(log, storage))
//})
//router.Get("/{alias}", redirect.New(log, storage))
//
//log.Info("starting server", slog.String("address", cfg.Address))
