package setting

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	file = "./config/config.yaml"
)

var (
	Cfg     *Conf
	ModPath = "~/.local/share/Terraria/tModLoader/Mods/"
	Proxy   = "http://localhost:7890"
)

type Conf struct {
	Version string `yaml:"version"`
	Http    struct {
		Port int `yaml:"port"`
	} `yaml:"http"`

	Log struct {
		Path       string `yaml:"path"`
		Level      string `yaml:"level"`
		MaxSize    int    `yaml:"maxsize"`
		MaxAge     int    `yaml:"maxage"`
		MaxBackups int    `yaml:"maxbackups"`
	} `yaml:"log"`
}

func init() {
	var err error
	Cfg, err = readConf()
	if err != nil {
		panic(err)
	}
}

func readConf() (*Conf, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var conf Conf

	err = yaml.NewDecoder(f).Decode(&conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
