package controller

import (
	"shoes/lib"
	"shoes/model"
)

type App struct {
	logging lib.Logging
	model   model.ModelApp
}

func New(db lib.Database) App {
	return App{model: model.New(db), logging: lib.LoadLogging()}
}
