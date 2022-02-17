package app

import (
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
	"io/ioutil"
	"net/http"
	"strings"
)

func (a *application) DeleteURLs(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(types.UserIDSessionKey).(types.Session)

	value, err := ioutil.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := strings.Split(strings.Trim(string(value), "[ ]"), ",")
	for i := 0; i < len(s); i++ {
		s[i] = strings.Trim(s[i], " \" \"")

		a.svc.DeleteURLS(session.UserID, s[i])
	}

	w.WriteHeader(http.StatusAccepted)
}
