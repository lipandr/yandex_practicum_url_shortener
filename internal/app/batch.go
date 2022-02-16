package app

import (
	"encoding/json"
	"fmt"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
	"io/ioutil"
	"net/http"
)

type BatchRequest struct {
	CID         string `json:"correlation_id"`
	OriginalURL string `json:"original_url"`
}

type BatchResponse struct {
	CID      string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}

func (a *application) Batch(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(types.UserIDSessionKey).(types.Session)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var batch []BatchRequest
	err = json.Unmarshal(body, &batch)
	if err != nil {
		fmt.Printf("Not Decoded: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp = make([]BatchResponse, len(batch))
	for i, v := range batch {
		s, err := a.svc.EncodeURL(session.UserID, v.OriginalURL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		out := fmt.Sprintf("%s/%s", a.cfg.BaseURL, s)
		resp[i] = BatchResponse{v.CID, out}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
