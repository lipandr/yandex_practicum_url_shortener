package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/app"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
	"log"
)

func main() {

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "File Storage Path")
	flag.StringVar(&cfg.DatabaseDsn, "d", cfg.DatabaseDsn, "Data base path string")
	flag.Parse()

	svc, err := service.NewService(cfg.FileStoragePath)
	if err != nil {
		log.Fatal("Can't start application")
	}
	urlApp := app.NewApp(cfg, svc)

	log.Fatal(urlApp.Run())
}
