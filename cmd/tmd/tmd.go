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
	p      *exec.Cmd
	in     io.WriteCloser
	ch     = make(chan bool)
	player []string
	//命令队列
	cmdInput = utils.NewQueue()
	//消息堆栈
	message = utils.NewStack()
	//是否开始
	started = false
)

func Start(start chan bool, file string) error {
	path := `./core/tModLoader/start-tModLoaderServer.sh`
	var proc *exec.Cmd
	if file == "" {
		proc = exec.Command("/bin/bash", "-c", path)
	} else {
		proc = exec.Command("/bin/bash", "-c", path, "-config", "./config/schemes/"+file+".txt")
	}
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
	go Command("start test")
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
	if !started {
		return "game not start!"
	}
	if cmd == "start test" {
		io.WriteString(in, cmd+"\n")
	} else if cmd == "exit" {
		p.Process.Kill()
		return "game has been closed."
	} else if cmd == "playing" {
		list := ""
		for i := 0; i < len(player); i++ {
			list = list + "$" + player[i]
		}
		return list
	}
	//写入命令
	fmt.Println(cmd, "start work")
	if _, err := io.WriteString(in, cmd+"\n"); err != nil {
		log.Println(err)
	}
	fmt.Println(cmd, "end of work")
	//wip 返回结果
	time.Sleep(500 * time.Millisecond)
	res := message.Pop()
	if strings.HasPrefix(res, ": ") {
		return strings.TrimPrefix(res, ": ")
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
				if strings.Contains(line, "has joined") {
					player = append(player, strings.TrimSuffix(line, "has joined."))
				}
				if strings.Contains(line, "has left") {
					for i := 0; i < len(player); i++ {
						if player[i] == strings.TrimSuffix(line, "has left.") {
							player = append(player[:i], player[i+1:]...)
						}
					}
				}
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
