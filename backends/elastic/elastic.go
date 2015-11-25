package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

type Elastic struct {
	cfg config.Elastic
}

func New(cfg config.Elastic) *Elastic {
	return &Elastic{
		cfg: cfg,
	}
}

func (e *Elastic) StoreDatagram(datagram config.Datagram) {
	url := fmt.Sprintf("http://%s%s/%s/event", e.cfg.Host, e.cfg.Port, e.cfg.IndexName)

	b, err := json.Marshal(datagram)
	if err != nil {
		log.Println("[ERROR] Couldn't marshal datagram into JSON")
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Println("[ERROR] Something went wrong with http.NewRequest")
		return
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
}
