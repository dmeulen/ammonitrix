package receiver

import (
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/backends/elastic"
	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

type Receiver struct {
	Config  *config.Config
	Elastic *elastic.Elastic
}

func NewReceiver(config *config.Config) (*Receiver, error) {
	r := &Receiver{
		Config: config,
	}

	return r, nil
}

var quit = make(chan error)

func (r *Receiver) StartListener() error {

	http.HandleFunc("/register", r.handleRegister)
	http.HandleFunc("/data", r.handleData)
	http.HandleFunc("/deregister", r.handleDeregister)
	go func() {
		err := http.ListenAndServe(r.Config.Listen.Port, nil)
		if err != nil {
			quit <- err
		}
	}()

	// wait for shutdown signal
	err := <-quit

	// trigger graceful shutdown
	log.Print("[INFO] Down: ", err)

	return nil
}

func (r *Receiver) ConnectElastic() error {
	el, err := elastic.NewElastic(r.Config)
	if err != nil {
		return err
	}

	r.Elastic = el
	return nil
}
