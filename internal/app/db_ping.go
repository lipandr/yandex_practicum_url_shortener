package app

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// DBPing handler проверяет соединение с базой данных.
// При успешной проверке хендлер возвращает HTTP-статус 200 OK,
// при неуспешной — 500 Internal Server Error.
func (a *application) DBPing(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", a.cfg.DatabaseDsn)
	if err != nil {
		panic(err)
	}

	defer func() {
		db.Close()
	}()

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
