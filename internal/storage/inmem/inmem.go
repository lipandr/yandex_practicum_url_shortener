package inmem

import "errors"

type Store struct {
	data map[string]string
}

var (
	keyNotFoundError     = errors.New("the key is not found")
	keyNotSpecifiedError = errors.New("the key is not specified")
)

func NewStorage() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", keyNotSpecifiedError
	}
	if value, ok := s.data[key]; ok {
		return value, nil
	}
	return "", keyNotFoundError
}

func (s *Store) Put(key, value string) error {
	if key == "" {
		return keyNotSpecifiedError
	}
	s.data[key] = value
	return nil
}
