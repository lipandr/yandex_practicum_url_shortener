package app

import (
	"encoding/json"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
	"net/http"
)

func (a *application) DeleteURLs(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(types.UserIDSessionKey).(types.Session)

	var IDs []string

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&IDs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(IDs); i++ {
		a.svc.DeleteURLS(session.UserID, IDs[i])
	}

	w.WriteHeader(http.StatusAccepted)
}
