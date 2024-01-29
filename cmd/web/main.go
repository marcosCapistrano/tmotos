package main

import (
	"fmt"
	"net/http"
	"os"
	"snippetbox/internal/config"
	"snippetbox/internal/db/models"
	"snippetbox/internal/db/models/repo/memory"
	"snippetbox/internal/logging"
)

type application struct {
	config     config.Configuration
	logger     logging.Logger
	repository models.Repository
}

func main() {
	config, err := config.Load("./config.json")
	if err != nil {
		panic(err)
	}

	logger := logging.NewDefaultLogger(config)
	repository := memory.NewMemoryRepo()
	repository.Seed()

	app := &application{
		config:     config,
		logger:     logger,
		repository: repository,
	}

	fmt.Println(os.Getenv("port"))
	fmt.Printf("Server listening on port %d\n", 4000)
	if err := http.ListenAndServe(":4000", app.routes()); err != nil {
		panic(err)
	}
}
