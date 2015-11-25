package receiver

// startListeners runs one or more listeners for the handler
import (
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

var quit = make(chan bool)

func StartListener(cfg *config.Config) {
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/data", handleData)
	http.HandleFunc("/deregister", handleDeregister)
	go http.ListenAndServe(cfg.Listen.Port, nil)

	// wait for shutdown signal
	<-quit

	// trigger graceful shutdown
	log.Print("[INFO] Down")
}
