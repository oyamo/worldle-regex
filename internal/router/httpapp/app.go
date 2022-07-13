package httpapp

import (
	"os"
	"os/signal"
	"reegle/config"
	v1 "reegle/internal/controller/http/v1"
	"reegle/pkg/dict"
	"reegle/pkg/server"
	"syscall"
)

func Run(cfg *config.Config) {
	db, err := dict.LoadDict(cfg.File.Path)
	if err != nil {
		cfg.Log.Fatal("dict load error ", err)
	} else {
		cfg.Log.Println("dict load success")
	}

	httpServer := server.New(v1.NewRouter(db, cfg), server.Port(cfg.Server.Port))

	// waiting signal
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interupt:
		cfg.Log.Println("router: received signal:", s.String())
	case err := <-httpServer.Notify():
		cfg.Log.Println("router: server error:", err)

	}

	// shutdown
	err = httpServer.Shutdown()
	if err != nil {
		cfg.Log.Println("router: server shutdown error:", err)
	}

}
