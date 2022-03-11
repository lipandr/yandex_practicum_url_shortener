package persistent

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

type Persistent struct {
	fileStoragePath string
}

func NewStorage(storagePath string) (*Persistent, error) {
	if storagePath == "" {
		return nil, errors.New("no path specified")
	}

	return &Persistent{
		fileStoragePath: storagePath,
	}, nil
}

func (s *Persistent) LoadURLsFromFile(m *inmem.Store) (err error) {
	file, err := os.OpenFile(s.fileStoragePath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			return
		}
		line := scanner.Text()
		data := strings.Split(line, " ")

		r := types.ShortenRecord{
			UserID: data[0],
			Key:    data[1],
			Value:  data[2],
		}

		if err := m.Put(r); err != nil {
			return err
		}
	}

	return
}

func (s *Persistent) StoreValue(r types.ShortenRecord) error {
	file, err := os.OpenFile(s.fileStoragePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	w := bufio.NewWriter(file)
	defer func() {
		_ = w.Flush()
	}()

	if _, err = w.WriteString(
		fmt.Sprintf("%s %s %s\n", r.UserID, r.Key, r.Value)); err != nil {
		return err
	}

	return nil
}
