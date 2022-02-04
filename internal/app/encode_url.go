package app

import (
	"fmt"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
	"io/ioutil"
	"net/http"
)

func (a *application) EncodeURL(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value("userID").(types.Session)
	value, err := ioutil.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	url := string(value)
	key, err := a.svc.EncodeURL(session.UserID, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(fmt.Sprintf("%s/%s", a.cfg.BaseURL, key)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
