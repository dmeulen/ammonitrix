package config

type Datagram struct {
	Timestamp  int         `json:"timestamp"`
	Name       string      `json:"name"`
	Hostname   string      `json:"hostname"`
	Tags       []string    `json:"tags"`
	Check_data interface{} `json:"check_data"`
}

type Register struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Interval int    `json:"interval"`
}

type Deregister struct {
	Name     string `json:"name"`
	Hostname string `json:"name"`
}
