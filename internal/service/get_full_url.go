package service

import "errors"

func (svc *service) GetFullURL(userID, key string) (string, error) {
	if svc.store[userID] != nil {
		res, err := svc.store[userID].Get(key)
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
