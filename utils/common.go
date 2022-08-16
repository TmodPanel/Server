package utils

import (
	"io"
	"net"
	"net/http"
)

func IpAddress() string {
	ip := "localhost"
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		addrs, _ := net.InterfaceAddrs()
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
				}
			}
		}
	} else {
		body, _ := io.ReadAll(resp.Body)
		ip = string(body)
	}
	return ip
}

func ErrToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
