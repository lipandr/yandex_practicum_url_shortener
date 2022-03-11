package service

import (
	"crypto/rand"
	"log"
	"math/big"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/lipandr/yandex_practicum_url_shortener/internal/storage/dao"
)

var benchUserID = uuid.NewString()

// URL encode benchmark
func BenchmarkService_EncodeURL(b *testing.B) {
	b.StopTimer()
	svc := dbConnection()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		url, _ := generateRandomURL()
		b.StartTimer()
		_, _ = svc.EncodeURL(benchUserID, url)
	}
}

func BenchmarkDBService_GetFullURL(b *testing.B) {
	b.StopTimer()
	svc := dbConnection()
	b.StartTimer()
	j := 1
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		id := strconv.Itoa(j)
		j++
		b.StartTimer()
		_, _ = svc.GetFullURL(id)
	}
}

// db connection helper method
func dbConnection() *dBService {
	db, err := dao.NewDB("postgres://localhost:5432/urlshorten?sslmode=disable")
	if err != nil {
		log.Fatal("Can't create DB connection:", err)
	}

	svc, err := NewDBService(db)
	if err != nil {
		log.Fatal("Can't allocate service:", err)
	}
	return svc
}

// helper method for encode URL benchmark
func generateRandomURL() (string, error) {
	const letters = "zxcvbnmasdfghjklqwertyuiop1234567890-"
	res := make([]byte, 10)
	for i := 0; i < 10; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		res[i] = letters[num.Int64()]
	}
	return string(res) + ".ru", nil
}
