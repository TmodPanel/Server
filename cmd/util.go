package main

import (
	"TSM-Server/utils"
	"fmt"
)

func main() {
	ip := utils.IpAddress()
	fmt.Println(ip)
}
