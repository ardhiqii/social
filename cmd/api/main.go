package main

import (
	"log"

	"github.com/ardhiqii/social/internal/env"
	"github.com/ardhiqii/social/internal/store"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
    log.Fatal("Error loading .env file")
  }

	cfg := config{
		addr: env.GetString("ADDR",":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store: store,
	}


	mux := app.mount()

	log.Fatal(app.run(mux))

}