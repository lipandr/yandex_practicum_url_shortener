package app

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	storage := inmem.NewStorage()
	svc := service.NewService(storage)
	app := NewApp(svc)

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vars := map[string]string{
				"key": tt.target,
			}
			request := httptest.NewRequest(tt.method, fmt.Sprintf("/%s", tt.target), bytes.NewReader([]byte(tt.body)))
			request = mux.SetURLVars(request, vars)

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
