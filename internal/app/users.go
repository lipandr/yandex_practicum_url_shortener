package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserJSON struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func (a *application) UserURLs(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	urls, err := a.svc.UsersURLs(userID)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var u []UserJSON

	for k, v := range urls {
		u = append(u, UserJSON{ShortURL: fmt.Sprintf("%s/%s", a.cfg.BaseURL, k), OriginalURL: v})
	}

	if err = json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
