package main

import (
	"log"

	"github.com/beparykamrul-dev/cgta/internal/config"
	"github.com/beparykamrul-dev/cgta/internal/httpapi"
)

func main() {
	cfg := config.Load()
	app := httpapi.New(cfg)
	if err := app.Listen(cfg.Address); err != nil {
		log.Fatal(err)
	}
}
