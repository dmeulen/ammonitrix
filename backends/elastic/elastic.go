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

func (e *Elastic) StoreDatagram(datagram config.Datagram) (*http.Response, error) {
	url := fmt.Sprintf("http://%s%s/%s/event", e.Config.Elastic.Host, e.Config.Elastic.Port, e.Config.Elastic.IndexName)

	b, err := json.Marshal(datagram)
	if err != nil {
		log.Println("[ERROR] Couldn't marshal datagram into JSON")
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Println("[ERROR] Something went wrong with http.NewRequest")
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

func (e *Elastic) StoreRegistration(reg config.Register) (*http.Response, error) {
	return nil, nil
}
