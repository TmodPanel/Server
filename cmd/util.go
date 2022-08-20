package main

import (
	"TSM-Server/internal/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var sz []service.Scheme
	data, _ := ioutil.ReadFile("./config/schemes/scheme.json")
	json.Unmarshal(data, &sz)
	fmt.Println(sz)
}
