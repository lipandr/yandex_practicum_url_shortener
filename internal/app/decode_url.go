package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (a *application) DecodeURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	url, err := a.svc.GetFullURL(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
