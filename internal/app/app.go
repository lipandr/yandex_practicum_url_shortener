package app

import (
	"github.com/gorilla/mux"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/app/middleware"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"

	"net/http"
)

type Application interface {
	Run() error
	EncodeURL(w http.ResponseWriter, r *http.Request)
	JSONEncodeURL(w http.ResponseWriter, r *http.Request)
	DecodeURL(w http.ResponseWriter, r *http.Request)
	DBPing(w http.ResponseWriter, r *http.Request)
	DefaultHandler(w http.ResponseWriter, r *http.Request)
}

type application struct {
	cfg config.Config
	svc service.Service
}

func NewApp(cfg config.Config, svc service.Service) Application {
	return &application{cfg: cfg, svc: svc}
}

func (a *application) Run() error {
	r := mux.NewRouter()

	r.Use(middleware.GzipMiddleware, middleware.AuthMiddleware)

	r.HandleFunc("/", a.EncodeURL).Methods(http.MethodPost)
	r.HandleFunc("/api/shorten/batch", a.Batch).Methods(http.MethodPost)
	r.HandleFunc("/api/shorten", a.JSONEncodeURL).Methods(http.MethodPost)
	r.HandleFunc("/ping", a.DBPing).Methods(http.MethodGet)
	r.HandleFunc("/user/urls", a.UserURLs).Methods(http.MethodGet)
	r.HandleFunc("/{key}", a.DecodeURL).Methods(http.MethodGet)
	r.HandleFunc("/", a.DefaultHandler)

	return http.ListenAndServe(a.cfg.ServerAddress, r)
}
