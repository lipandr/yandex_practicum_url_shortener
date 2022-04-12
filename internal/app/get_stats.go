package app

import (
	"encoding/json"
	"net/http"
)

type APIStatsJSONResponse struct {
	URLs  int `json:"urls"`  // Количество сокращённых URL в сервисе
	Users int `json:"users"` // Количество пользователей в сервисе
}

func (a *application) GetStats(w http.ResponseWriter, _ *http.Request) {
	urls, users, err := a.svc.GetStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res := APIStatsJSONResponse{
		URLs:  urls,
		Users: users,
	}
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
