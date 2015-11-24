package receiver

import (
	"encoding/json"
	"fmt"
	"io"
)

func validateDataRequest(body io.Reader) (bool, Datagram) {
	decoder := json.NewDecoder(body)
	var d Datagram
	err := decoder.Decode(&d)
	if err != nil {
		fmt.Println(err)
		return false, d
	}
	return true, d
}
