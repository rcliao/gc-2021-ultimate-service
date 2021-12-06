package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var build = "develop"

func main() {
	log.Println("starting service TEST. build:", build)
	defer log.Println("shutdown")

	// to receive the shutdown signal from Kubernetes
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("starting shutdown")
}
