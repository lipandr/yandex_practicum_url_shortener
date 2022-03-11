package app

import "net/http"

// DefaultHandler handler заглушка используемая по умолчанию.
// Возвращает HTTP-статус 400 BadRequest.
func (a *application) DefaultHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}
