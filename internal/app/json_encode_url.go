package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ApiJSONRequest struct {
	Url string `json:"url"`
}

type ApiJSONResponse struct {
	Result string `json:"result"`
}

func (r ApiJSONRequest) Validate() error {
	if r.Url == "" {
		return errors.New("incorrect JSON url")
	}
	return nil
}

func (a *application) JSONEncodeURL(w http.ResponseWriter, r *http.Request) {
	var req ApiJSONRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var url = req.Url

	key, err := a.svc.EncodeURL(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res := ApiJSONResponse{
		Result: fmt.Sprintf("%s/%s", a.cfg.BaseURL, key),
	}
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
