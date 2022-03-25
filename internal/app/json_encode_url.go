package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

// APIJSONRequest структура описывающая json-формат запроса.
type APIJSONRequest struct {
	URL string `json:"url"`
}

// APIJSONResponse структура описывающая json-формат ответа.
type APIJSONResponse struct {
	Result string `json:"result"`
}

// Validate helper-метод для валидации запросов.
func (r APIJSONRequest) Validate() error {
	if r.URL == "" {
		return errors.New("incorrect JSON url")
	}

	return nil
}

// JSONEncodeURL handler принимающий в теле запроса данные в json-формате.
// При успешном сокращении URL возвращает HTTP-статус 201 Created.
// При повторном сокращении URL, сервис возвратит HTTP-статус 409 Conflict.
func (a *application) JSONEncodeURL(w http.ResponseWriter, r *http.Request) {
	var req APIJSONRequest
	session := r.Context().Value(types.UserIDSessionKey).(types.Session)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := http.StatusCreated
	var url = req.URL

	key, err := a.svc.EncodeURL(session.UserID, url)
	if err != nil {
		if errors.Is(err, types.ErrKeyExists) {
			status = http.StatusConflict
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	res := APIJSONResponse{
		Result: fmt.Sprintf("%s/%s", a.cfg.BaseURL, key),
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
