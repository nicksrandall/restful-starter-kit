package app

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/ExpansiveWorlds/instrumentedsql"
	opentracingsql "github.com/ExpansiveWorlds/instrumentedsql/opentracing"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

type contextKey = int

var associationKey contextKey = 0

func InitDB() (*sqlx.DB, error) {
	// connect to the database
	sqlLogger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		Logger(ctx).Sugar().Infow(msg, keyvals)
	})

	_ = sqlLogger

	sql.Register(
		"instrumented-postgres",
		instrumentedsql.WrapDriver(
			&pq.Driver{},
			instrumentedsql.WithTracer(opentracingsql.NewTracer()),
			// instrumentedsql.WithLogger(sqlLogger),
		),
	)
	db, err := sql.Open("instrumented-postgres", Config.DSN)
	if err != nil {
		return nil, err
	}
	dbx := sqlx.NewDb(db, "postgres")

	// handle db migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return nil, err
	}

	Logger(nil).Sugar().Infof("applied %d migrations", n)

	return dbx, nil
}

func DBMiddleware(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), associationKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func DBFromContext(ctx context.Context) *sqlx.DB {
	db, _ := ctx.Value(associationKey).(*sqlx.DB)
	return db
}
