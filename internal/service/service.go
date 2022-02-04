package service

import (
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/persistent"
)

type Service interface {
	GetFullURL(userID, key string) (string, error)
	EncodeURL(userID, url string) (string, error)
	UsersURLs(userID string) (map[string]string, error)
}

type service struct {
	store      map[string]*inmem.Store
	persistent *persistent.Persistent
}

func NewService(storagePath string) (*service, error) {
	p, err := persistent.NewStorage(storagePath)
	if err != nil {
		return nil, err
	}

	ss := make(map[string]*inmem.Store)
	err = p.LoadURLsFromFile(ss)
	if err != nil {
		return nil, err
	}

	return &service{
		store:      ss,
		persistent: p,
	}, nil
}
