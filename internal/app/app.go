package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
	"net/http"
)

type Application interface {
	Run() error
	EncodeUrl(w http.ResponseWriter, r *http.Request)
	DecodeUrl(w http.ResponseWriter, r *http.Request)
	DefaultHandler(w http.ResponseWriter, r *http.Request)
}

type application struct {
	svc service.Service
}

func NewApp(svc service.Service) *application {
	return &application{svc: svc}
}

func (a *application) Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/", a.EncodeUrl).Methods("POST")
	r.HandleFunc("/{key}", a.DecodeUrl).Methods("GET")
	r.HandleFunc("/", a.DefaultHandler)

	return http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), r)
}
