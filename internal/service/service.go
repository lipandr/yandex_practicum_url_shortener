package service

import (
	store "github.com/lipandr/yandex_practicum_url_shortener/internal/storage"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/persistent"
)

type Service interface {
	GetFullURL(key string) (string, error)
	EncodeURL(url string) (string, error)
}

type service struct {
	store      store.Repository
	persistent *persistent.Persistent
}

func NewService(storagePath string) (*service, error) {
	p, err := persistent.NewStorage(storagePath)
	if err != nil {
		return nil, err
	}
	im := inmem.NewStorage()
	err = p.LoadURLsFromFile(im)
	if err != nil {
		return nil, err
	}
	return &service{
		store:      im,
		persistent: p,
	}, nil
}
