package main

import (
	"sample_rest/app"
	"sample_rest/config"
)

func main() {
	configuration := config.GetConfig()

	application := &app.App{}
	application.Initialize(configuration)
	application.Run(":3000")
}
