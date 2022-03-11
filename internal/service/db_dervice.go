package service

import (
	"fmt"
	"strconv"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

func (svc *dBService) EncodeURL(userID, url string) (string, error) {
	var id int64
	var isDeleted bool

	row := svc.db.QueryRow("SELECT url_id, is_deleted FROM url WHERE original = ($1);", url)
	err := row.Scan(&id, &isDeleted)
	if err != nil {
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

	if isDeleted {
		q := "UPDATE url SET created_by = ($1), is_deleted=false WHERE url_id = ($2);"
		_, err := svc.db.Exec(q, userID, id)
		if err != nil {
			return "", fmt.Errorf("addURL: %v", err)
		}
		return strconv.FormatInt(id, 10), nil
	}

	return strconv.FormatInt(id, 10), types.ErrKeyExists
}

func (svc *dBService) GetFullURL(key string) (string, error) {
	var orig string
	var isDeleted bool

	row := svc.db.QueryRow("SELECT original, is_deleted FROM url WHERE url_id = ($1);", key)
	err := row.Scan(&orig, &isDeleted)
	if err != nil {
		return "", fmt.Errorf("original URL Error: %v", err)
	}

	if isDeleted {
		return "", types.ErrKeyDeleted
	}

	return orig, nil
}

func (svc dBService) UsersURLs(userID string) map[string]string {
	urls := make(map[string]string)

	q := "SELECT url_id, original from url where created_by = ($1) AND is_deleted=false;"
	rows, err := svc.db.Query(q, userID)
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

func (svc *dBService) DeleteURLS(userID string, url string) {
	urlID, err := strconv.Atoi(url)
	if err != nil {
		return
	}

	q := "UPDATE url SET is_deleted=true WHERE url_id = ($1) AND created_by = ($2);"
	_, err = svc.db.Exec(q, urlID, userID)
	if err != nil {
		fmt.Printf("URL exsists: %v\n", err)
	}
}
