package config

import (
	"encoding/json"
	"flag"
	"github.com/caarlos0/env/v6"
	"io/ioutil"
	"log"
)

// Config структура описывающая переменные окружения и их значения по умолчанию.
// Флаг -a, отвечающий за адрес запуска HTTP-сервера (переменная SERVER_ADDRESS);
// флаг -b, отвечающий за базовый адрес результирующего сокращённого URL (переменная BASE_URL);
// флаг -f, отвечающий за путь до файла с сокращёнными URL (переменная FILE_STORAGE_PATH);
// флаг -d, отвечающий за путь подключения к базе данных с сокращенными URL (переменная DATABASE_DSN);
// флаг -s, отвечающий за запуск сервера приложения в режиме HTTPS (переменная ENABLE_HTTPS);
// флаг -c, отвечающий возможность конфигурации приложения с помощью файла в формате JSON.
type Config struct {
	ServerAddress   string `json:"server_address" env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `json:"base_url" env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `json:"file_storage_path" env:"FILE_STORAGE_PATH" envDefault:""`
	DatabaseDsn     string `json:"database_dsn" env:"DATABASE_DSN" envDefault:"postgres://localhost:5432/urlshorten?sslmode=disable"`
	EnableHTTPS     bool   `json:"enable_https" env:"ENABLE_HTTPS" envDefault:"false"`
	Config          string `env:"CONFIG"`
}

func InitConfig() Config {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "Server address")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "Base URL")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "File Storage Path")
	flag.StringVar(&cfg.DatabaseDsn, "d", cfg.DatabaseDsn, "Data base path string")
	flag.BoolVar(&cfg.EnableHTTPS, "s", cfg.EnableHTTPS, "Enable HTTPS server mode")
	flag.StringVar(&cfg.Config, "c", cfg.Config, "JSON config file")
	flag.Parse()

	// Read and parse JSON file if flag -c with value exists
	jsonFileData, err := ioutil.ReadFile(cfg.Config)
	if err != nil {
		return cfg
	}
	var jsonCfg Config
	if err = json.Unmarshal(jsonFileData, &jsonCfg); err != nil {
		return cfg
	}
	if cfg.ServerAddress == "" {
		cfg.ServerAddress = jsonCfg.ServerAddress
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = jsonCfg.BaseURL
	}
	if cfg.FileStoragePath == "" {
		cfg.FileStoragePath = jsonCfg.FileStoragePath
	}
	if cfg.DatabaseDsn == "" {
		cfg.DatabaseDsn = jsonCfg.DatabaseDsn
	}
	if !cfg.EnableHTTPS || jsonCfg.EnableHTTPS {
		cfg.EnableHTTPS = jsonCfg.EnableHTTPS
	}

	return cfg
}
