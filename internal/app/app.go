package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/app/middleware"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
)

// Application интерфейс описывающий методы приложения.
type Application interface {
	EncodeURL(w http.ResponseWriter, r *http.Request)
	JSONEncodeURL(w http.ResponseWriter, r *http.Request)
	DecodeURL(w http.ResponseWriter, r *http.Request)
	DBPing(w http.ResponseWriter, r *http.Request)
	DefaultHandler(w http.ResponseWriter, r *http.Request)
	AppShutdown(w http.ResponseWriter, r *http.Request)
	GetStats(w http.ResponseWriter, r *http.Request)
}

// Структура Dependency injection приложения.
type application struct {
	http.Server
	cfg         config.Config
	svc         service.Service
	shutdownReq chan bool
	reqCount    uint32
}

// NewApp метод-конструктор приложения
func NewApp(cfg config.Config, svc service.Service) *application {
	a := &application{
		Server: http.Server{
			Addr:         cfg.ServerAddress,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		cfg:         cfg,
		svc:         svc,
		shutdownReq: make(chan bool),
	}

	r := mux.NewRouter()
	r.Use(
		middleware.GzipMiddleware,
		middleware.AuthMiddleware,
		middleware.TrustedSubnetMiddleware(a.cfg.TrustedSubnet),
	)

	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

	r.HandleFunc("/", a.EncodeURL).Methods(http.MethodPost)
	r.HandleFunc("/api/internal/stats", a.GetStats).Methods(http.MethodGet)
	r.HandleFunc("/api/admin/shutdown", a.AppShutdown).Methods(http.MethodGet)
	r.HandleFunc("/api/shorten/batch", a.Batch).Methods(http.MethodPost)
	r.HandleFunc("/api/shorten", a.JSONEncodeURL).Methods(http.MethodPost)
	r.HandleFunc("/ping", a.DBPing).Methods(http.MethodGet)
	r.HandleFunc("/api/user/urls", a.UserURLs).Methods(http.MethodGet)
	r.HandleFunc("/api/user/urls", a.DeleteURLs).Methods(http.MethodDelete)
	r.HandleFunc("/{key}", a.DecodeURL).Methods(http.MethodGet)
	r.HandleFunc("/", a.DefaultHandler)

	a.Handler = r

	return a
}
