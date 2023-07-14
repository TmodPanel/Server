package server

import (
	"encoding/json"
	"os"
)

type Map struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Seed     string `json:"seed"`
	WordSize string `json:"wordSize"`
	Mode     string `json:"mode"`
	Type     string `json:"type"`
}

var mapJSONPath = "./data/map/maps.json"

func GetMapInfo() ([]Map, error) {
	f, err := os.ReadFile(mapJSONPath)
	if err != nil {
		return nil, err
	}

	var list []Map
	err = json.Unmarshal(f, &list)

	return list, err
}
