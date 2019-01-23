package servers

import (
	"fmt"
)

type Restful struct {
	Host string
	Port int16
	Url string
}

func NewRestful(host string, port int16) *Restful {
	url := fmt.Sprintf("http://%s:%d", host, port)
	return &Restful{host, port, url}
}