package service

import (
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/persistent"
)

type Service interface {
	GetFullURL(userID, key string) (string, error)
	EncodeURL(userID, url string) (string, error)
	UsersURLs(userID string) map[string]string
}

type service struct {
	inMem      *inmem.Store
	persistent *persistent.Persistent
}

func NewService(storagePath string) (*service, error) {
	inMem := inmem.NewStorage()

	//a := make([]*ShortenRecord, 0)

	var p *persistent.Persistent

	if storagePath != "" {
		p, err := persistent.NewStorage(storagePath)
		if err != nil {
			return nil, err
		}
		err = p.LoadURLsFromFile(inMem)
		if err != nil {
			return nil, err
		}
	}

	return &service{
		inMem:      inMem,
		persistent: p,
	}, nil
}
