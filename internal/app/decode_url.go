package app

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
	"net/http"
)

func (a *application) DecodeURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	url, err := a.svc.GetFullURL(key)
	if err != nil {
		if errors.Is(err, types.ErrKeyDeleted) {
			w.WriteHeader(http.StatusGone)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
