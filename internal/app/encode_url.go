package app

import (
	"fmt"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"io/ioutil"
	"net/http"
)

func (a *application) EncodeURL(w http.ResponseWriter, r *http.Request) {
	value, err := ioutil.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	url := string(value)
	key, err := a.svc.EncodeURL(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(fmt.Sprintf("http://%s:%d/%s", config.Host, config.Port, key)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
