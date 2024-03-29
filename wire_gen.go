// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/MichalMichalak/wire-playground/process"
	"github.com/MichalMichalak/wire-playground/store"
)

// Injectors from wire.go:

func initialize() process.Worker {
	options := newOptions()
	postgres := store.NewPostgres(options)
	worker := process.NewWorker(postgres)
	return worker
}
