// Учебный сервис по сокращению URL yandex-practicum.
package main

import (
	"flag"
	"fmt"
	"log"

	_ "net/http/pprof"

	"github.com/caarlos0/env/v6"

	"github.com/lipandr/yandex_practicum_url_shortener/internal/app"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/config"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/service"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/dao"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Println("Build version:", buildVersion)
	fmt.Println("Build date:", buildDate)
	fmt.Println("Build commit:", buildCommit)

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "File Storage Path")
	flag.StringVar(&cfg.DatabaseDsn, "d", cfg.DatabaseDsn, "Data base path string")
	flag.BoolVar(&cfg.EnableHTTPS, "s", cfg.EnableHTTPS, "Enable HTTPS server mode")
	flag.Parse()

	db, err := dao.NewDB(cfg.DatabaseDsn)
	if err != nil {
		log.Fatal("Can't start application:", err)
	}

	svc, err := service.NewDBService(db)
	if err != nil {
		log.Fatal("Can't start application:", err)
	}

	urlApp := app.NewApp(cfg, svc)
	log.Fatal(urlApp.Run())
}
