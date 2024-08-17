package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/brutalzinn/go-reverse-proxy/config"
	"github.com/brutalzinn/go-reverse-proxy/proxy"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatalf("Error loading config: %v", err)
	}
	///chan blocking signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	for _, route := range cfg.Routes {
		go proxy.StartProxy(route)
	}
	<-done
}
