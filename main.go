package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//Start a tcp server
	tcp := &server{Address: ":53", Net: "tcp"}
	tcp.startServer()

	//start a UDP server
	udp := &server{Address: ":53", Net: "udp"}
	udp.startServer()

	//Creating a channel for signal propogation
	sig := make(chan os.Signal)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Printf("Caught Signal %+v", <-sig)
		log.Println("... shutting down...")
		os.Exit(0)
	}()

	//wait
	select {}
}
