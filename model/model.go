package model

import (
	"context"
	"shoes/lib"
	"time"
)

const defaultTimeout = 60 * time.Second

type ModelApp struct {
	db lib.Database
}

func New(db lib.Database) ModelApp {
	return ModelApp{db: db}
}

func defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
