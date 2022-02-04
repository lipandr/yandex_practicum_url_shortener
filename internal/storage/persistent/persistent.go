package persistent

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"os"
	"strings"
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

func (s *Persistent) LoadURLsFromFile(ss map[string]*inmem.Store) (err error) {
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

		if ss[data[2]] == nil {
			ss[data[2]] = inmem.NewStorage()
		}

		if err := ss[data[2]].Put(data[0], data[1]); err != nil {
			return err
		}
	}
	return
}

func (s *Persistent) StoreValue(userId, key, value string) error {
	file, err := os.OpenFile(s.fileStoragePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	writer := bufio.NewWriter(file)
	defer func() {
		_ = writer.Flush()
	}()

	if _, err = writer.WriteString(fmt.Sprintf("%s %s %s\n", key, value, userId)); err != nil {
		return err
	}

	return nil
}
