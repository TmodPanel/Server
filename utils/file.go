package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	ModPath = "~/.local/share/Terraria/tModLoader/Mods/"
)

func read(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var lines []string
	for {
		line, _ := r.ReadString('\n')
		if line == "" {
			break
		}
		lines = append(lines, strings.Trim(line, "\r\n"))
	}
	return lines, nil
}
func write(file string, lines []string, filter string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, line := range lines {
		if !strings.Contains(line, filter) {
			fmt.Fprintf(f, "%s\n", line)
		}
	}
	return nil
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
func RemoveFromBanList(name string) error {
	lines, err := read("./config/banlist.txt")
	if err != nil {
		return err
	}
	return write("./config/banlist.txt", lines, name)
}
func ReadUserList() ([]string, error) {
	lines, err := read("./config/userList.txt")
	if err != nil {
		return nil, err
	}
	return lines, err
}
func GetModInfo() (map[string]bool, error) {
	files, err := os.ReadDir(ModPath)
	if err != nil {
		return nil, err
	}
	all := make(map[string]bool)
	for i := 0; i < len(files); i++ {
		if strings.Contains(files[i].Name(), ".tmod") {
			name := strings.Split(files[i].Name(), ".tmod")[0]
			all[name] = false
		}
	}
	mods, err := read(ModPath + "enabled.json")
	valid := regexp.MustCompile("[a-zA-Z]")
	for i := 1; i < len(mods)-1; i++ {
		mod := strings.Join(valid.FindAllString(mods[i], -1), "")
		all[mod] = true
	}
	return all, err
}
func RemoveFromEnable(name string) error {
	lines, err := read(ModPath + "enabled.json")
	if err != nil {
		return err
	}
	err = write(ModPath+"enabled.json", lines, name)
	return err
}
func EnableMod(name string) error {
	if err := RemoveFromEnable(name); err != nil {
		return err
	}
	lines, err := read(ModPath + "enabled.json")
	if err != nil {
		return err
	}
	n := len(lines)
	lines[n-1] = fmt.Sprintf(`  "%s",`, name)
	lines = append(lines, "]")
	return write(ModPath+"enabled.json", lines, "*")
}
func DelMod(name string) error {
	err := os.Remove(ModPath + name + ".tmod")
	return err
}
