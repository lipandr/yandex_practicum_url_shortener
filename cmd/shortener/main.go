package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/app"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/persistent"
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
	flag.Parse()

	fmt.Println(cfg.FileStoragePath)
	rep := persistent.NewStorage(cfg.FileStoragePath)
	svc := service.NewService(rep)
	urlApp := app.NewApp(cfg, svc)

	log.Fatal(urlApp.Run())
}
