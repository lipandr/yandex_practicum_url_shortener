module github.com/lipandr/yandex_practicum_url_shortener

go 1.15

replace github.com/lipandr/yandex_practicum_url_shortener => ./

require (
	github.com/caarlos0/env/v6 v6.9.1
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.4
	github.com/xo/dburl v0.9.0
)
