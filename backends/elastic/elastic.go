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
	Config *config.Config
}

func NewElastic(config *config.Config) (*Elastic, error) {
	e := &Elastic{
		Config: config,
	}

	return e, nil
}

func (e *Elastic) StoreDatagram(datagram config.Datagram) {
	log.Println("before url")
	url := fmt.Sprintf("http://%s%s/%s/event", e.Config.Elastic.Host, e.Config.Elastic.Port, e.Config.Elastic.IndexName)
	log.Println(url)

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
	log.Println("here")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
}
