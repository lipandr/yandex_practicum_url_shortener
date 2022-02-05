package app

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func (a *application) DBPing(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", a.cfg.DatabaseDsn)
	if err != nil {
		panic(err)
	}

	defer func() {
		db.Close()
	}()

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
