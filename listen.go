package main

// startListeners runs one or more listeners for the handler
import (
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

var quit = make(chan bool)

func startListener(listen []config.Listen, h http.Handler) {
	for _, l := range listen {
		go http.ListenAndServe(l.Addr, h)
	}

	// wait for shutdown signal
	<-quit

	// trigger graceful shutdown
	log.Print("[INFO] Down")
}
