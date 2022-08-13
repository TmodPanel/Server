package tmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	path      = `/home/user/Downloads/tModLoader/start-tModLoaderServer.sh`
	proc      = exec.Command("/bin/bash", "-c", path, "-config", "server.config")
	stdout, _ = proc.StdoutPipe()
	stdin, _  = proc.StdinPipe()
	stderr, _ = proc.StderrPipe()
	ch        = make(chan bool)
	cmdline   = NewQueue()
)

func init() {
	if err := proc.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		os.Exit(1)
	}

}

func Start(start chan bool) {
	go asyncLog(stdout)
	go asyncLog(stderr)
	go Command("")

	ok := <-ch
	start <- ok
	log.Println("Server started")

	if err := proc.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......", err.Error())
	}
}

func Command(cmd string) string {
	fmt.Println(cmd, "start work")
	_, err := io.WriteString(stdin, cmd+"\n")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cmd, "end of work")
	res := ""
	for {
		res = cmdline.Pop()
		if strings.Contains(res, ":") || res == "" {
			continue
		} else {
			break
		}
	}
	return res
}

func asyncLog(reader io.ReadCloser) {
	cache := ""
	buf := make([]byte, 1024, 1024)
	started := false

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
			cache = oSlice[len(oSlice)-1]
			line = line + cache
			fmt.Println(line)
			if started {
				cmdline.Push(line)
			}
			if strings.Contains(line, "Server started") {
				started = true
				ch <- true
			}
		}
	}
}

type Queue struct {
	list []string
}

func NewQueue() *Queue {
	list := make([]string, 0)
	return &Queue{list: list}
}
func (q *Queue) Push(data string) {
	q.list = append(q.list, data)
}

func (q *Queue) Pop() string {
	if len(q.list) == 0 {
		return ""
	}
	res := q.list[0]
	q.list = q.list[1:]
	return res
}
