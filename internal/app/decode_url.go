package app

import (
	"github.com/gorilla/mux"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
	"net/http"
)

func (a *application) DecodeURL(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(types.UserIdSessionKey).(types.Session)
	vars := mux.Vars(r)
	key := vars["key"]

	url, err := a.svc.GetFullURL(session.UserID, key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
