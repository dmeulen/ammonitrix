package receiver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

func (r *Receiver) handleRegister(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "registered!\n")
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
	r.Elastic.StoreDatagram(datagram)
	err := json.NewEncoder(w).Encode(datagram)
	if err != nil {
		http.Error(w, "Failed to encode datagram", 500)
	}
}
