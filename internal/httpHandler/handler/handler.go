package handler

import (
	"github.com/DeMarDeXis/VProj/internal/httpHandler/handler/mw/logger"
	"github.com/DeMarDeXis/VProj/internal/service"
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

	router.Route("/app", func(r chi.Router) {
		r.Use(h.userIdentity)
		r.Route("/nhl", func(nhlRout chi.Router) {
			nhlRout.Get("/teams", h.getTeams)
			//r.Get("/teams/{id}", h.getTeam)
			//r.Get("/players", h.getPlayers)
			//r.Get("/players/{id}", h.getPlayer)
			//r.Get("/games", h.getGames)
			//r.Get("/games/{id}", h.getGame)
			//r.Get("/standings", h.getStandings)
			//r.Get("/standings/{id}", h.getStandingsByTeam)
			nhlRout.Get("/schedule", h.getSchedule)
			nhlRout.Get("/schedule/last/{id}", h.getLastSchedule)
			//r.Get("/schedule/{id}", h.getScheduleByTeam)
			//r.Get("/stats", h.getStats)
			//r.Get("/stats/{id}", h.getStatsByPlayer)
		})
		//r.Route("/nba", func(nbaRout chi.Router) {
		//	nbaRout.Get("/teams", h.getTeams)
		//	//r.Get("/teams/{id}", h.getTeam)
		//	//r.Get("/players", h.getPlayers)
		//	//r.Get("/players/{id}", h.getPlayer)
		//	//r.Get("/games", h.getGames)
		//	//r.Get("/games/{id}", h.getGame)
		//	//r.Get("/standings", h.getStandings)
		//	//r.Get("/standings/{id}", h.getStandingsByTeam)
		//	//r.Get("/schedule", h.getSchedule)
		//	//r.Get("/schedule/{id}", h.getScheduleByTeam)
		//	//r.Get("/stats", h.getStats)
		//	//r.Get("/stats/{id}", h.getStatsByPlayer)
		//})
		//r.Route("/nfl", func(nflRout chi.Router) {
		//	nflRout.Get("/teams", h.getTeams)
		//	//r.Get("/teams/{id}", h.getTeam)
		//	//r.Get("/players", h.getPlayers)
		//	//r.Get("/players/{id}", h.getPlayer)
		//	//r.Get("/games", h.getGames)
		//	//r.Get("/games/{id}", h.getGame)
		//	//r.Get("/standings", h.getStandings)
		//	//r.Get("/standings/{id}", h.getStandingsByTeam)
		//	//r.Get("/schedule", h.getSchedule)
		//	//r.Get("/schedule/{id}", h.getScheduleByTeam)
		//	//r.Get("/stats", h.getStats)
		//	//r.Get("/stats/{id}", h.getStatsByPlayer)
		//})
		//r.Post("/tasks", h.createList)
		//r.Get("/tasks", h.getAllLists)
		//r.Get("/tasks/{id}", h.getListByID)
		//r.Put("/tasks/{id}", h.updateList)
		//r.Delete("/tasks/{id}", h.deleteList)
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
