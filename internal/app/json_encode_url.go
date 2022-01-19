package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type APIJSONRequest struct {
	URL string `json:"url"`
}

type APIJSONResponse struct {
	Result string `json:"result"`
}

func (r APIJSONRequest) Validate() error {
	if r.URL == "" {
		return errors.New("incorrect JSON url")
	}
	return nil
}

func (a *application) JSONEncodeURL(w http.ResponseWriter, r *http.Request) {
	var req APIJSONRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var url = req.URL

	key, err := a.svc.EncodeURL(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := APIJSONResponse{
		Result: fmt.Sprintf("%s/%s", a.cfg.BaseURL, key),
	}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}