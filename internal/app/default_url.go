package app

import "net/http"

func (a *application) DefaultHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}
