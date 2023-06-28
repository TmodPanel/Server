package utils

import "strings"

func WriteServerConf(args []string, file string) error {
	err := write("./config/schemes/"+file+".txt", args, "^_^")
	return err
}

func ReadServerConfig(conf string) (string, error) {
	lines, err := read("./config/serverconfig.txt")
	if err != nil {
		return "", err
	}
	for _, line := range lines {
		if strings.Contains(line, conf+"=") && !strings.HasPrefix(line, "#") {
			res := strings.TrimLeft(line, conf+"=")
			return res, nil
		}
	}
	return "Unable to get", nil
}
