package service

import (
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/persistent"
)

type Service interface {
	GetFullURL(key string) (string, error)
	EncodeURL(userID, url string) (string, error)
	UsersURLs(userID string) map[string]string
}

type service struct {
	inMem      *inmem.Store
	persistent *persistent.Persistent
}

func NewService(storagePath string) (*service, error) {
	var p *persistent.Persistent
	inMem := inmem.NewStorage()

	if storagePath != "" {
		n, err := persistent.NewStorage(storagePath)
		if err != nil {
			return nil, err
		}
		err = n.LoadURLsFromFile(inMem)
		if err != nil {
			return nil, err
		}
		// TODO Check here
		p = n
	}

	return &service{
		inMem:      inMem,
		persistent: p,
	}, nil
}
