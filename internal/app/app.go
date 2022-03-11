package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/app/middleware"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
)

// Application интерфейс описывающий методы приложения.
type Application interface {
	Run() error
	EncodeURL(w http.ResponseWriter, r *http.Request)
	JSONEncodeURL(w http.ResponseWriter, r *http.Request)
	DecodeURL(w http.ResponseWriter, r *http.Request)
	DBPing(w http.ResponseWriter, r *http.Request)
	DefaultHandler(w http.ResponseWriter, r *http.Request)
}

// Структура Dependency injection приложения.
type application struct {
	cfg config.Config
	svc service.Service
}

// NewApp метод-конструктор приложения
func NewApp(cfg config.Config, svc service.Service) Application {
	return &application{cfg: cfg, svc: svc}
}

// Run метод запуска приложения
func (a *application) Run() error {
	r := mux.NewRouter()

	r.Use(middleware.GzipMiddleware, middleware.AuthMiddleware)

	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	r.HandleFunc("/", a.EncodeURL).Methods(http.MethodPost)
	r.HandleFunc("/api/shorten/batch", a.Batch).Methods(http.MethodPost)
	r.HandleFunc("/api/shorten", a.JSONEncodeURL).Methods(http.MethodPost)
	r.HandleFunc("/ping", a.DBPing).Methods(http.MethodGet)
	r.HandleFunc("/api/user/urls", a.UserURLs).Methods(http.MethodGet)
	r.HandleFunc("/api/user/urls", a.DeleteURLs).Methods(http.MethodDelete)
	r.HandleFunc("/{key}", a.DecodeURL).Methods(http.MethodGet)

	r.HandleFunc("/", a.DefaultHandler)

	return http.ListenAndServe(a.cfg.ServerAddress, r)
}
