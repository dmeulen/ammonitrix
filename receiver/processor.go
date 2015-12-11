package receiver

import (
	"encoding/json"
	"io"
	"log"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

func (r *Receiver) validateDataRequest(body io.Reader) (bool, config.Datagram) {
	decoder := json.NewDecoder(body)
	var d config.Datagram
	err := decoder.Decode(&d)
	if err != nil {
		log.Println(err)
		return false, d
	}
	return true, d
}

func (r *Receiver) validateRegisterRequest(body io.Reader) (bool, config.Register) {
	decoder := json.NewDecoder(body)
	var reg config.Register
	err := decoder.Decode(&reg)
	if err != nil {
		log.Println(err)
		return false, reg
	}
	return true, reg
}

func (r *Receiver) validateDeregisterRequest(body io.Reader) (bool, config.Deregister) {
	decoder := json.NewDecoder(body)
	var dereg config.Deregister
	err := decoder.Decode(&dereg)
	if err != nil {
		log.Println(err)
		return false, dereg
	}
	return true, dereg
}
