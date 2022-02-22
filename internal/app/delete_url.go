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

	numWorkers := 3
	numJobs := len(IDs)
	jobs := make(chan job, numJobs)
	results := make(chan job, numJobs)

	for w := 0; w < numWorkers; w++ {
		go a.worker(jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		item := job{
			userID: session.UserID,
			urlID:  IDs[j],
		}
		jobs <- item
	}
	close(jobs)

	for a := 0; a < numJobs; a++ {
		<-results
	}

	w.WriteHeader(http.StatusAccepted)
}

func (a *application) worker(jobs <-chan job, results chan<- job) {
	for j := range jobs {
		a.svc.DeleteURLS(j.userID, j.urlID)
		results <- j
	}
}

type job struct {
	userID string
	urlID  string
}
