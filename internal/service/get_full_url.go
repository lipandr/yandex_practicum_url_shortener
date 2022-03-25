package service

import "errors"

// GetFullURL сервис получения полного URL из памяти.
func (svc *service) GetFullURL(key string) (string, error) {
	if svc.inMem != nil {
		res, err := svc.inMem.Get(key)
		if err != nil {
			return "", err
		}
		return res, nil
	}

	return "", errors.New("not found")
}

// UsersURLs сервис получения сокращенных пользователем URL из памяти.
func (svc *service) UsersURLs(userID string) map[string]string {
	if svc.inMem != nil {
		return svc.inMem.GetAllUserKeys(userID)
	}

	return map[string]string{}
}

// DeleteURLS метод-заглушка для реализации интерфейса Service.
func (svc *service) DeleteURLS(userID string, url string) {
}
