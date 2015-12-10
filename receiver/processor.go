package receiver

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

func (r *Receiver) validateDataRequest(body io.Reader) (bool, config.Datagram) {
	decoder := json.NewDecoder(body)
	var d config.Datagram
	err := decoder.Decode(&d)
	if err != nil {
		fmt.Println(err)
		return false, d
	}
	return true, d
}
