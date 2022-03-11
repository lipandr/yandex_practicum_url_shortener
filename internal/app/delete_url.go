package app

import (
	"encoding/json"
	"net/http"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

// Структура job используемая worker
type job struct {
	userID string
	urlID  string
}

// DeleteURLs handler удаляет сохраненные пользователем URL
func (a *application) DeleteURLs(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(types.UserIDSessionKey).(types.Session)

	var IDs []string

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&IDs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numWorkers := 3
	numJobs := len(IDs)

	jobs := make(chan job, numJobs)
	defer close(jobs)

	for w := 0; w < numWorkers; w++ {
		go a.worker(jobs)
	}

	for j := 0; j < numJobs; j++ {
		item := job{
			userID: session.UserID,
			urlID:  IDs[j],
		}
		jobs <- item
	}

	w.WriteHeader(http.StatusAccepted)
}

// Helper-метод worker используемый для удаления URL.
func (a *application) worker(jobs <-chan job) {
	for j := range jobs {
		a.svc.DeleteURLS(j.userID, j.urlID)
	}
}
