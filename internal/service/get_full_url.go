package service

func (svc *service) GetFullURL(key string) (string, error) {
	res, err := svc.store.Get(key)
	if err != nil {
		return "", err
	}
	return res, nil
}
