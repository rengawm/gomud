package main

import (
	"log"
	"os"
	"os/signal"
)

type GameManager struct {
	Net *NetManager
}

func main() {
	gameManager := &GameManager{
		Net: NewNetManager(2000),
	}

	go func() { log.Fatal(gameManager.Net.Start()) }()
	waitForExit()
}

// waitForExit blocks until an interrupt or kill signal is received by the program.
func waitForExit() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	log.Printf("Got %v signal", sig)
	return
}
