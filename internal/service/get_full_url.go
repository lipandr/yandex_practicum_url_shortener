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

func (svc *service) UsersURLs(userID string) (map[string]string, error) {
	res, err := svc.store[userID].GetAllKeys()
	if err != nil {
		return nil, err
	}

	return res, nil
}
