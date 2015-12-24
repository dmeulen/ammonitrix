package receiver

import (
	"io"
	"log"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

func (r *Receiver) handleRegister(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(w, "Unsupported method", 405)
		return
	}
	var valid bool
	var reg config.Register
	if valid, reg = r.validateRegisterRequest(req.Body); valid != true {
		http.Error(w, "Invalid request received", 500)
		return
	}
	log.Printf("[DEBUG] Received a registration request: %s", reg)
	resp, err := r.Elastic.StoreRegistration(reg)
	if err != nil {
		http.Error(w, "Failed to store registration", 500)
	}

	w.WriteHeader(resp.StatusCode)
}

func (r *Receiver) handleDeregister(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "deregistered!\n")
}

func (r *Receiver) handleData(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(w, "Unsupported method", 405)
		return
	}
	var valid bool
	var datagram config.Datagram
	if valid, datagram = r.validateDataRequest(req.Body); valid != true {
		http.Error(w, "Invalid request received", 500)
		return
	}
	log.Printf("[DEBUG] Received a valid datagram: %s", datagram)
	resp, err := r.Elastic.StoreDatagram(datagram)
	if err != nil {
		http.Error(w, "Failed to store datagram", 500)
	}

	w.WriteHeader(resp.StatusCode)
}
