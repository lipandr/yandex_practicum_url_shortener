package service

import (
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

// EncodeURL сервис сохранения URL в память и в файл на диск
func (svc *service) EncodeURL(userID, url string) (string, error) {

	//if svc.store[userID] == nil {
	//	svc.store[userID] = inmem.NewStorage()
	//}
	hash := svc.generateHash()

	r := types.ShortenRecord{
		UserID: userID,
		Key:    hash,
		Value:  url,
	}

	if err := svc.inMem.Put(r); err != nil {
		return "", err
	}

	if svc.persistent != nil {
		err := svc.persistent.StoreValue(r)
		if err != nil {
			return "", err
		}
	}

	return hash, nil
}

// Helper-метод generateHash
func (svc *service) generateHash() string {
	return svc.inMem.GetCurrentSeq()
}
