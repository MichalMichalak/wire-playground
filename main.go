package main

import "github.com/MichalMichalak/wire-playground/store"

func main() {
	worker := initialize()
	result := worker.Work(6789)
	println("Worker output", result)
}

func newOptions() store.Options {
	return store.Options{
		Host: "host-1",
		Port: 12345,
		DB:   "mydb",
	}
}
