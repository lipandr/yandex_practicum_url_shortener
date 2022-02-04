package service

import "errors"

func (svc *service) GetFullURL(userID, key string) (string, error) {
	if us, ok := svc.store[userID]; ok {
		res, err := us.Get(key)
		if err != nil {
			return "", err
		}
		return res, nil
	}
	return "", errors.New("not found")
}

func (svc *service) UsersURLs(userID string) map[string]string {
	if svc.store[userID] != nil {
		return svc.store[userID].GetAllKeys()
	}
	return map[string]string{}
}
