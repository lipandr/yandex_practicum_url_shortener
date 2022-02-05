package service

import (
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
)

func (svc *service) EncodeURL(userID, url string) (string, error) {

	if svc.store[userID] == nil {
		svc.store[userID] = inmem.NewStorage()
	}
	hash := svc.generateHash(userID)
	err := svc.store[userID].Put(hash, url)
	if err != nil {
		return "", err
	}

	if svc.persistent != nil {
		err := svc.persistent.StoreValue(userID, hash, url)
		if err != nil {
			return "", err
		}
	}

	return hash, nil
}

func (svc *service) generateHash(userID string) string {
	return svc.store[userID].GetCurrentSeq()
}
