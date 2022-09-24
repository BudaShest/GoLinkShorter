package main

import (
	"shorter/config"
	"shorter/internal/app"
	"shorter/pkg/helpers"
)

//TODO удалить всё fmt. Юзать только для дебага

func main() {
	cfg, err := config.LoadConfig("./../../config.yaml")
	helpers.FailOnError(err, "Configuration error")

	application := app.GetAppInstance()
	err = application.Init(cfg)

	helpers.FailOnError(err, "Initialization error")

	application.Run()
}
