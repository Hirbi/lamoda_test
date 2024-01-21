package main

import (
	"app/config"
	"app/server"
	"log"
)

func main() {
	err := config.InitConifg()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp(config.Config.Server.Port)
	err = app.Run()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
