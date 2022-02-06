package service

import (
	"fmt"
	"strconv"
)

func (svc dBService) GetFullURL(key string) (string, error) {
	var orig string

	row := svc.db.QueryRow("SELECT original FROM url WHERE url_id = ($1);", key)
	err := row.Scan(&orig)
	if err != nil {
		return "", fmt.Errorf("original URL Error: %v", err)
	}

	return orig, nil

}

func (svc dBService) EncodeURL(userID, url string) (string, error) {
	var id int64

	_, err := svc.db.Exec("INSERT INTO url (original, created_by) VALUES ($1, $2);", url, userID)
	if err != nil {
		return "", fmt.Errorf("addURL: %v", err)
	}

	row := svc.db.QueryRow("SELECT url_id FROM url WHERE original = ($1);", url)
	err = row.Scan(&id)
	if err != nil {
		return "", fmt.Errorf("URL ID Error: %v", err)
	}

	return strconv.FormatInt(id, 10), nil
}

func (svc dBService) UsersURLs(userID string) map[string]string {
	urls := make(map[string]string)

	rows, err := svc.db.Query("SELECT url_id, original from url where created_by = ($1);", userID)
	if err != nil {
		return nil
	}

	defer func() {
		rows.Close()
	}()

	for rows.Next() {
		var u, o string

		err = rows.Scan(&u, &o)
		if err != nil {
			return nil
		}
		urls[u] = o

	}

	err = rows.Err()
	if err != nil {
		return nil
	}

	return urls
}
