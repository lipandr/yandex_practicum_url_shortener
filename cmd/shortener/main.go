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
	done := make(chan bool)

	go func() {
		if cfg.EnableHTTPS {
			// Generate private key (.key)
			// openssl genrsa -out server.key 2048
			// Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
			// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
			err := urlApp.ListenAndServeTLS("server.crt", "server.key")
			if err != nil {
				log.Printf("Listen and serve: %v", err)
			}
			done <- true
		}

		err := urlApp.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
		done <- true
	}()

	urlApp.WaitShutdown()

	<-done
	log.Printf("DONE!")
}
