package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

var (
	ModPath string
)

func init() {
	u, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	if u.Username == "root" {
		ModPath = fmt.Sprintf("/root/.local/share/Terraria/tModLoader/Mods/")
	} else {
		ModPath = fmt.Sprintf("/home/%s/.local/share/Terraria/tModLoader/Mods/", u.Username)
	}
}
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
	var all map[string]bool
	for i := 0; i < len(files); i++ {
		if strings.Contains(files[i].Name(), ".tmod") {
			all[files[i].Name()] = false
		}
	}
	lines, err := read("./config/enabled.json")
	for i := 1; i < len(lines)-1; i++ {
		str := strings.Trim(lines[i], ",")
		n := len(str)
		all[str[3:n-1]] = true
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
