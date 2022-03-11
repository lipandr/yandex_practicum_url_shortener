package app

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/types"
)

func TestHandlers(t *testing.T) {
	uid := uuid.NewString()

	cfg := config.Config{
		ServerAddress: "localhost:8080",
		BaseURL:       "http://localhost:8080",
	}

	svc, err := service.NewService("")
	if err != nil {
		t.Fatal(err)
	}

	app := NewApp(cfg, svc)

	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name    string
		handler func(w http.ResponseWriter, r *http.Request)
		method  string
		target  string
		body    string
		want    want
	}{
		{
			name:    "positive encode test #1",
			handler: app.EncodeURL,
			method:  http.MethodPost,
			target:  "",
			body:    "https://google.com",
			want: want{
				code:     201,
				response: `http://localhost:8080/1`,
			},
		},
		{
			name:    "positive decode test #1",
			handler: app.DecodeURL,
			method:  http.MethodGet,
			target:  "1",
			body:    "",
			want: want{
				code:        307,
				response:    "<a href=\"https://google.com\">Temporary Redirect</a>.\n\n",
				contentType: "text/html; charset=utf-8",
			},
		},
		{
			name:    "negative decode test #1",
			handler: app.DecodeURL,
			method:  http.MethodGet,
			target:  "7",
			body:    "",
			want: want{
				code: 400,
			},
		},
		{
			name:    "positive encode test #2",
			handler: app.EncodeURL,
			method:  http.MethodPost,
			target:  "",
			body:    "https://yandex.ru",
			want: want{
				code:     201,
				response: `http://localhost:8080/2`,
			},
		},
		{
			name:    "positive decode test #2",
			handler: app.DecodeURL,
			method:  http.MethodGet,
			target:  "2",
			body:    "",
			want: want{
				code:        307,
				response:    "<a href=\"https://yandex.ru\">Temporary Redirect</a>.\n\n",
				contentType: "text/html; charset=utf-8",
			},
		},
		{
			name:    "positive JSON test #1",
			handler: app.JSONEncodeURL,
			method:  http.MethodPost,
			target:  "/api/shorten",
			body:    "{\"url\": \"https://google.com\"}",
			want: want{
				code:        201,
				response:    "{\"result\":\"http://localhost:8080/3\"}\n",
				contentType: "application/json",
			},
		},
		{
			name:    "negative JSON test #1",
			handler: app.JSONEncodeURL,
			method:  http.MethodPost,
			target:  "/api/shorten",
			body:    "{\"url\": \"\"}",
			want: want{
				code:        400,
				response:    "incorrect JSON url\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vars := map[string]string{
				"key": tt.target,
			}
			request := httptest.NewRequest(tt.method, fmt.Sprintf("/%s", tt.target), bytes.NewReader([]byte(tt.body)))
			request = mux.SetURLVars(request, vars)

			ctx := context.WithValue(request.Context(), types.UserIDSessionKey, types.Session{UserID: uid})
			request = request.WithContext(ctx)

			w := httptest.NewRecorder()
			h := http.HandlerFunc(tt.handler)

			h.ServeHTTP(w, request)
			res := w.Result()

			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

			defer res.Body.Close()

			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			if string(resBody) != tt.want.response {
				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
			}

			if res.Header.Get("Content-Type") != tt.want.contentType {
				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
			}
		})
	}
}
