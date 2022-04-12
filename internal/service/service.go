package service

import (
	"database/sql"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/persistent"
)

// Service интерфейс сервиса приложения.
type Service interface {
	GetFullURL(key string) (string, error)
	EncodeURL(userID, url string) (string, error)
	UsersURLs(userID string) map[string]string
	DeleteURLS(userID string, url string)
	GetStats() (int, int, error)
}

// Структура dBService описывающая интерфейс подключения к БД.
type dBService struct {
	db *sql.DB
}

// NewDBService метод-конструктор для dBService.
func NewDBService(db *sql.DB) (*dBService, error) {
	return &dBService{
		db: db,
	}, nil
}

// Структура service описывающая хранение информации в памяти и на жестком диске.
type service struct {
	inMem      *inmem.Store
	persistent *persistent.Persistent
}

// NewService метод-конструктор для service.
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
		p = n
	}

	return &service{
		inMem:      inMem,
		persistent: p,
	}, nil
}
