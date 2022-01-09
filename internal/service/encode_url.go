package service

import "strconv"

var i int

func (svc *service) EncodeURL(url string) (string, error) {
	hash := svc.generateHash()
	err := svc.store.Put(hash, url)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (svc *service) generateHash() string {
	i++
	return strconv.Itoa(i)
}
