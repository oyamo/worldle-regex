package main

import (
	"log"
	"reegle/config"
	server "reegle/internal/router/httpapp"
)

func main() {
	// start server
	cfg, err := config.LoadConfig(".env.example", config.DEV)
	if err != nil {
		log.Fatalln(err)
	}

	server.Run(cfg)

}
