package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var version = "0.0.0"

type MySignal struct {
	message string
}

func (s MySignal) String() string {
	return s.message
}

func (s MySignal) Signal() {}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Println("version:", version)
		return
	}

	log.Println("[info] Start")
	trapSingals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}

	sigCh := make(chan os.Signal, 1)

	time.AfterFunc(5*time.Second, func() {
		sigCh <- MySignal{"time out"}
	})

	signal.Notify(sigCh, trapSingals...)

	sig := <-sigCh
	switch s := sig.(type) {
	case syscall.Signal:
		log.Printf("[info] Got signal: %s(%d)", s, s)
	case MySignal:
		log.Printf("[info] %s", s)
	}
}
