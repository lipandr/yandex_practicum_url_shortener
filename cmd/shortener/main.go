package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//Напишите сервис для сокращения длинных URL.
//Требования:
//Сервер должен быть доступен по адресу: http://localhost:8080.
//Сервер должен предоставлять два эндпоинта: POST / и GET /{id}.
//Эндпоинт POST / принимает в теле запроса строку URL для сокращения
// и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.

//Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL
// и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.

//Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400.
const host = "localhost:8080"
var id int
var urlsStore = make(map[string]string)
// key: shorten URL
// value: URL

func PostHandler(w http.ResponseWriter, r *http.Request) {

	value, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	url := string(value)
	key := generateShortenUrl(url)

	err = putKey(key, url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(key))

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]


	if value, ok := urlsStore[key]; ok {
		http.Redirect(w, r, value, http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {
	urlsStore["ya.ru"] = "http://yandex.ru"

	r := mux.NewRouter()

	r.HandleFunc("/", PostHandler).Methods("POST")
	r.HandleFunc("/{key}", GetHandler).Methods("GET")
	r.HandleFunc("/", DefaultHandler)

	log.Fatal(http.ListenAndServe(host, r))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func putKey(key string, value string) error {
	urlsStore[key] = value
	return nil
}

func generateShortenUrl(s string) string {

	u := strconv.Itoa(id)
	id++
	urlsStore[u] = s
	return fmt.Sprintf("http://%s/%s", host, u)
}
