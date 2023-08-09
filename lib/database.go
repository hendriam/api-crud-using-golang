package lib

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const defaultTimeout = 2 * time.Second

type Database struct {
	*sql.DB
}

func LoadDatabase() (Database, error) {
	config := LoadConfig()
	logging := LoadLogging()
	dsn := config.Database.MySQL.Dsn
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logging.Error().Msgf("[Database] failed connected")
		return Database{}, err
	}

	ctx, cancel := defaultContext()
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		logging.Error().Msgf("[Database] failed connected")
		return Database{}, err
	}

	logging.Info().Msgf("[Database] success connected")
	return Database{db}, nil
}

func defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
