package main

import (
	"sample_rest/app"
	"sample_rest/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
