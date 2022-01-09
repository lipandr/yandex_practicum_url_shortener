package inmem

import "errors"

type Store struct {
	data map[string]string
}

var (
	errKeyNotFound     = errors.New("the key is not found")
	errKeyNotSpecified = errors.New("the key is not specified")
)

func NewStorage() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", errKeyNotSpecified
	}
	if value, ok := s.data[key]; ok {
		return value, nil
	}
	return "", errKeyNotFound
}

func (s *Store) Put(key, value string) error {
	if key == "" {
		return errKeyNotSpecified
	}
	s.data[key] = value
	return nil
}
