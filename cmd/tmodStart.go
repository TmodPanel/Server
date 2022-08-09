package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	path := `./start-tModLoaderServer.sh`
	proc := exec.Command("/bin/bash", "-c", path, "-config", "server.config")
	stdout, _ := proc.StdoutPipe()
	stderr, _ := proc.StderrPipe()
	//stdin, _ := proc.StdinPipe()

	if err := proc.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		os.Exit(1)
	}

	go asyncLog(stdout)
	go asyncLog(stderr)

	if err := proc.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......", err.Error())
	}

}

func asyncLog(reader io.ReadCloser) error {
	cache := ""
	buf := make([]byte, 1024, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "closed") {
				err = nil
			}
		}
		if num > 0 {
			oByte := buf[:num]
			oSlice := strings.Split(string(oByte), "\n")
			line := strings.Join(oSlice[:len(oSlice)-1], "\n")
			fmt.Println(cache, line)
			cache = oSlice[len(oSlice)-1]
		}
	}
	return nil
}
