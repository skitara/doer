package main

import (
	"log"

	"github.com/skitara/doer"
	"github.com/skitara/doer/cmd/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(doer.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
