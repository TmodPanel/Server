package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

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
