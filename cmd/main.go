package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	// h :=api.mount()
	// api.run(h) or,

	if err := api.run(api.mount()); err != nil {
		log.Println("Server has failed to start")
		os.Exit(1)
	}
}
