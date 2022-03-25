package inmem

import (
	"errors"
	"strconv"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

var (
	errKeyNotFound     = errors.New("the key is not found")
	errKeyNotSpecified = errors.New("the key is not specified")
)

type Store struct {
	globalStore map[string]string
	userStore   map[string][]string
}

func NewStorage() *Store {
	return &Store{
		globalStore: make(map[string]string),
		userStore:   make(map[string][]string),
	}
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", errKeyNotSpecified
	}

	if value, ok := s.globalStore[key]; ok {
		return value, nil
	}

	return "", errKeyNotFound
}

func (s *Store) GetAllUserKeys(uid string) map[string]string {
	k := make(map[string]string)

	if _, ok := s.userStore[uid]; ok {
		for _, seqList := range s.userStore[uid] {
			if seq, ok := s.globalStore[seqList]; ok {
				k[seqList] = seq
			}
		}
	}

	return k
}

func (s *Store) Put(r types.ShortenRecord) error {
	if r.Key == "" {
		return errKeyNotSpecified
	}

	s.globalStore[r.Key] = r.Value
	s.userStore[r.UserID] = append(s.userStore[r.UserID], r.Key)

	return nil
}

func (s *Store) GetCurrentSeq() string {
	return strconv.Itoa(len(s.globalStore) + 1)
}
