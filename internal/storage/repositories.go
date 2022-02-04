package storage

type Repository interface {
	Put(key string, value string) error
	Get(key string) (string, error)
	GetAllKeys() (map[string]string, error)
	GetCurrentSeq() string
}
