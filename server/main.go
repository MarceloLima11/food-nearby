package main

import (
	"log"

	"github.com/MarceloLima11/food-nearby/server/config"
	"github.com/MarceloLima11/food-nearby/server/middleware"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("Config initialization error: %v", err)
	}

	middleware.Initialize()
}
