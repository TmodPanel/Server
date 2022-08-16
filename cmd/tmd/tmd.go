package tmd

import (
	"TSM-Server/utils"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	p  *exec.Cmd
	in io.WriteCloser
	ch = make(chan bool)
	//命令队列
	cmdInput = utils.NewQueue()
	//消息堆栈
	message = utils.NewStack()
	//是否开始
	started = false
)

func Start(start chan bool) error {
	path := `/home/user/Downloads/tModLoader/start-tModLoaderServer.sh`
	proc := exec.Command("/bin/bash", "-c", path, "-config", "server.config")
	p = proc
	stdout, _ := proc.StdoutPipe()
	stdin, _ := proc.StdinPipe()
	in = stdin
	stderr, _ := proc.StderrPipe()
	if err := proc.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		return err
		os.Exit(1)
	}
	go asyncLog(stdout)
	go asyncLog(stderr)
	go Command("")
	ok := <-ch
	start <- ok
	log.Println("Server started")

	if err := proc.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......", err.Error())
		return err
	}
	return nil
}

func Command(cmd string) string {
	cmdInput.Push(cmd)
	fmt.Println(cmd, "start work")
	_, err := io.WriteString(in, cmd+"\n")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cmd, "end of work")
	time.Sleep(1200 * time.Millisecond)
	res := message.Pop()
	if res == "" {
		return "game not start!"
	}
	if strings.HasPrefix(res, ": ") {
		return strings.TrimPrefix(res, ": ")
	}
	if cmd == "exit" && started {
		p.Process.Kill()
	}
	return res
}

func asyncLog(reader io.ReadCloser) {
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
			cache = oSlice[len(oSlice)-1]
			line = line + cache
			fmt.Println(line)
			if started {
				message.Push(line)
			}
			if strings.Contains(line, "Server started") {
				started = true
				ch <- true
			}
		}
	}
}

func CheckStart() bool {
	b, _ := exec.Command("ps", "-ef").Output()
	if strings.Contains(string(b), "start-tModLoaderServer.sh") && started {
		return true
	}
	return false
}
