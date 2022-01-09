package main

//
//import (
//	"io/ioutil"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestDefaultHandler(t *testing.T) {
//	type want struct {
//		code        int
//		response    string
//		contentType string
//	}
//	tests := []struct {
//		name string
//		want want
//	}{
//		{
//			name: "400 code test",
//			want: want{
//				code: 400,
//				response: "",
//				contentType: "",
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			request := httptest.NewRequest(http.MethodGet, "/status", nil)
//
//			// создаём новый Recorder
//			w := httptest.NewRecorder()
//			// определяем хендлер
//			h := http.HandlerFunc(DefaultHandler)
//			// запускаем сервер
//			h.ServeHTTP(w, request)
//			res := w.Result()
//
//			// проверяем код ответа
//			if res.StatusCode != tt.want.code {
//				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
//			}
//
//			// получаем и проверяем тело запроса
//			defer res.Body.Close()
//			resBody, err := ioutil.ReadAll(res.Body)
//			if err != nil {
//				t.Fatal(err)
//			}
//			if string(resBody) != tt.want.response {
//				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
//			}
//
//			// заголовок ответа
//			if res.Header.Get("Content-Type") != tt.want.contentType {
//				t.Errorf("Expected Content-Type %s, got %s", tt.want.contentType, res.Header.Get("Content-Type"))
//			}
//		})
//	}
//}
//
//func TestGetHandler(t *testing.T) {
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
//
//func TestPostHandler(t *testing.T) {
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
//
//func Test_putKey(t *testing.T) {
//	type args struct {
//		key   string
//		value string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			name: "No error test",
//			args: args{
//				key: "1",
//				value: "http://ya.ru",
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := putKey(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
//				t.Errorf("putKey() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
