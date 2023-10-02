package main

import (
	"github.com/MarceloLima11/food-nearby/server/config"
	"github.com/MarceloLima11/food-nearby/server/middleware"
	"github.com/MarceloLima11/food-nearby/server/utils"
)

func main() {
	err := config.Init()
	utils.IfErrThrowFatalf(err, "Config initialization error:")

	middleware.Initialize()
}
