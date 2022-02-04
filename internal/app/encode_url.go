package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (a *application) EncodeURL(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	value, err := ioutil.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	url := string(value)
	key, err := a.svc.EncodeURL(userID, url)
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
