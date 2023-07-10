package example

import (
	"fmt"
	"net/http"
)

const dnsIP = "192.168.0.1"

var (
	ip   = "192.168.12.42"
	port = 3333
)

func DialPlace() {
	http.Get(fmt.Sprintf("http://%s:%d", ip, port))
	someIP := "192.168.0.1"
	http.Get(fmt.Sprintf("http://%s:%d", someIP, port))
}
