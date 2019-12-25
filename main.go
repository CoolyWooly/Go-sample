package main

import (
	"os"
	"sample_rest/app"
	"sample_rest/config"
)

func main() {
	configuration := config.GetConfig()

	application := &app.App{}
	application.Initialize(configuration)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	application.Run(port)

	//application.Run(":5000")
}
