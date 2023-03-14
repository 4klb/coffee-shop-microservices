package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/4klb/coffeetime/dish/config"
)

var (
	dbConnOnce sync.Once
	db         *sql.DB
)

func GetDB(ctx context.Context) *sql.DB {
	dbConnOnce.Do(func() {
		cfg := config.GetConfig()
		psql := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DB)

		db, err := sql.Open("postgres", psql)
		if err != nil {
			log.Println(err)
			return
		}

		db.SetMaxIdleConns(20)

		if err := db.PingContext(ctx); err != nil {
			log.Println(err)
			return
		}
	})

	return db
}
