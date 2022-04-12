package app

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

// DecodeURL handler возвращает полный, сохраненный URL.
// При успешном выполнении возвращает HTTP-ответ 307 Temporary redirect.
// Если URL ранее не сохранялся ни одним пользователем,
// вернет HTTP-статус 400 BadRequest.
// Если URL был удален пользователем, будет возвращен HTTP-статус 410 Gone.
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
