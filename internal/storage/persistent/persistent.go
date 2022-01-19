package persistent

import (
	"bufio"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"log"
	"os"
	"strings"
)

type Persistent struct {
	FileStoragePath string
	inMemory        *inmem.Store
}

func NewStorage(storagePath string) *Persistent {
	inMemory := inmem.NewStorage()

	if err := loadURLsFromFile(inMemory, storagePath); err != nil {
		log.Fatal(err)
	}
	return &Persistent{
		FileStoragePath: storagePath,
		inMemory:        inMemory,
	}
}

func loadURLsFromFile(inMemory *inmem.Store, fileStoragePath string) (err error) {
	if fileStoragePath == "" {
		return
	}

	file, err := os.OpenFile(fileStoragePath, os.O_RDONLY|os.O_CREATE, 0777)
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

		if err := inMemory.Put(data[0], data[1]); err != nil {
			return err
		}
	}
	return
}

func (s *Persistent) Get(key string) (string, error) {
	value, err := s.inMemory.Get(key)

	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *Persistent) Put(key, value string) error {
	if err := s.inMemory.Put(key, value); err != nil {
		return err
	}

	if s.FileStoragePath != "" {
		file, err := os.OpenFile(s.FileStoragePath, os.O_WRONLY|os.O_APPEND, 0777)
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

		if _, err = writer.WriteString(key + " " + value + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (s *Persistent) GetCurrentSeq() string {
	return s.inMemory.GetCurrentSeq()
}
