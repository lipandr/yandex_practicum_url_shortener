// Учебный сервис по сокращению URL yandex-practicum.
package main

import (
	"fmt"
	"log"

	_ "net/http/pprof"

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

	cfg := config.InitConfig()

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
