package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

// UserJSON структура описывающая json-формат запроса
type UserJSON struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// UserURLs handler возвращает список URL сокращенных пользователем
// Если у пользователя нет сокращенных URL, возвращает HTTP-статус 204 NoContent.
func (a *application) UserURLs(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(types.UserIDSessionKey).(types.Session)

	urls := a.svc.UsersURLs(session.UserID)
	if len(urls) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var u []UserJSON

	for k, v := range urls {
		u = append(u, UserJSON{ShortURL: fmt.Sprintf("%s/%s", a.cfg.BaseURL, k), OriginalURL: v})
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
