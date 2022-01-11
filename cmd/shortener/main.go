package main

import (
	"log"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/app"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/inmem"
)

func main() {
	rep := inmem.NewStorage()
	svc := service.NewService(rep)
	urlApp := app.NewApp(svc)

	log.Fatal(urlApp.Run())
}
