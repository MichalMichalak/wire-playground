//+build wireinject

package main

import (
	"github.com/MichalMichalak/wire-playground/process"
	"github.com/MichalMichalak/wire-playground/store"
	"github.com/google/wire"
)

func initialize() process.Worker {
	wire.Build(process.NewWorker, wire.Bind(new(process.DataWarehouse), new(store.Postgres)), store.NewPostgres, newOptions)
	return process.Worker{}
}
