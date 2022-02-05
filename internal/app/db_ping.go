package app

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/xo/dburl"
	"net/http"
	"time"
)

func (a *application) DBPing(w http.ResponseWriter, r *http.Request) {
	db, err := dburl.Open(a.cfg.DatabaseDsn)
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
