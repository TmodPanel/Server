package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	path := `/home/user/Downloads/tModLoader/start-tModLoaderServer.sh`
	proc := exec.Command("/bin/bash", "-c", path, "-config", "server.config")
	stdout, _ := proc.StdoutPipe()
	stdin, _ := proc.StdinPipe()
	//stderr, _ := proc.StderrPipe()

	if err := proc.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		os.Exit(1)
	}

	//flag := make(chan bool)
	go func() {
		buf := make([]byte, 1024, 1024)
		for {
			num, err := stdout.Read(buf)
			if err != nil {
				if err == io.EOF || strings.Contains(err.Error(), "closed") {
					err = nil
				}
			}
			if num > 0 {
				oByte := buf[:num]
				oSlice := strings.Split(string(oByte), "\n")
				line := strings.Join(oSlice[:len(oSlice)-1], "\n")
				fmt.Println(line)
				if strings.Contains(line, "Server started") {

				}
			}
		}
	}()
	go func() {
		for {
			time.Sleep(3 * time.Second)
			stdin.Write([]byte("help"))
		}
	}()

	//go func() {
	//	for {
	//		time.Sleep(10 * time.Second)
	//		stdin.Write([]byte("seed"))
	//	}
	//
	//}()

	if err := proc.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......", err.Error())
	}

}
