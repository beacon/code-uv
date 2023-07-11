package example

import (
	"fmt"
	"net/http"
)

const dnsIP = "192.168.0.1"

var (
	ip       = "192.168.12.42"
	port     = 3333
	maskIP   = "255.255.255.0"
	loopback = "127.0.0.1"
)

func DialPlace() {
	http.Get(fmt.Sprintf("http://%s:%d", ip, port))
	someIP := "192.168.0.22"
	ipCommentedOut := "192.168.0.33" // IGNORE_CODE_SCAN: this ip is here for testing
	http.Get(fmt.Sprintf("http://%s:%d", ipCommentedOut, port))
	http.Get(fmt.Sprintf("http://%s:%d", someIP, port))
	http.Get(fmt.Sprintf("http://%s:%d", maskIP, port))
	http.Get(fmt.Sprintf("http://%s:%d", loopback, port))
}
