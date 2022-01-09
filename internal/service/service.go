package service

import store "github.com/lipandr/yandex_practicum_url_shortener/internal/storage"

type Service interface {
	GetFullURL(key string) (string, error)
	EncodeURL(url string) (string, error)
}

type service struct {
	store store.Repository
}

func NewService(storage store.Repository) *service {
	return &service{
		store: storage,
	}
}
