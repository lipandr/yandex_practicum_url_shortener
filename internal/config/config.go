package config

// Config структура описывающая переменные окружения и их значения по умолчанию.
// Флаг -a, отвечающий за адрес запуска HTTP-сервера (переменная SERVER_ADDRESS);
// флаг -b, отвечающий за базовый адрес результирующего сокращённого URL (переменная BASE_URL);
// флаг -f, отвечающий за путь до файла с сокращёнными URL (переменная FILE_STORAGE_PATH);
// флаг -d, отвечающий за путь подключения к базе данных с сокращенными URL (переменная DATABASE_DSN);
// флаг -s, отвечающий за запуск сервера приложения в режиме HTTPS (переменная ENABLE_HTTPS).
type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:"test.txt"`
	DatabaseDsn     string `env:"DATABASE_DSN" envDefault:"postgres://localhost:5432/urlshorten?sslmode=disable"`
	EnableHTTPS     bool   `env:"ENABLE_HTTPS" envDefault:"false"`
}
