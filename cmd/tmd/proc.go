package tmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type OutputMessage struct {
	OutputBuf  [100]string // message buffer
	outputPos  int         // current position in buffer
	outputSize int         // current size of buffer
}

type PlayerStatus struct {
	players map[string]time.Time
	count   int
}

type TModProc struct {
	proc    *exec.Cmd
	inPipe  io.WriteCloser
	outPipe io.ReadCloser
	message *OutputMessage
	ps      *PlayerStatus
	outChan chan string
	errChan chan error
	IsOpen  bool
}

func NewTModLoader(config string) (*TModProc, error) {
	path := `./core/tModLoader/start-tModLoaderServer.sh`
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("Error getting file info: %s......", err.Error())
		return nil, err
	}
	if fileInfo.Mode().Perm()&0111 == 0 {
		// if the file is not executable, make it executable
		err = os.Chmod(path, fileInfo.Mode().Perm()|0100)
		if err != nil {
			log.Printf("Error getting file info: %s......", err.Error())
			return nil, err
		}
	}

	proc := exec.Command(path, "-config", config)

	stdout, err := proc.StdoutPipe()
	if err != nil {
		log.Printf("Error getting stdout pipe: %s......", err.Error())
		return nil, err
	}

	stdin, err := proc.StdinPipe()
	if err != nil {
		log.Printf("Error getting stdin pipe: %s......", err.Error())
		return nil, err
	}

	return &TModProc{
		proc:    proc,
		inPipe:  stdin,
		outPipe: stdout,
		outChan: make(chan string),
		errChan: make(chan error),
		IsOpen:  false,
		message: &OutputMessage{
			OutputBuf:  [100]string{},
			outputPos:  0,
			outputSize: 0,
		},
		ps: &PlayerStatus{
			players: make(map[string]time.Time),
			count:   0,
		},
	}, nil
}

func (t *TModProc) Start() error {
	if err := t.proc.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		return err
	}

	// send a newline to the process to start it
	_, _ = t.inPipe.Write([]byte("\n"))

	go ServerMonitor(t)

	if err := t.proc.Wait(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		return err
	}
	return nil
}

func ServerMonitor(t *TModProc) {
	scanner := bufio.NewScanner(t.outPipe)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, SERVER_STARTED) {
			t.IsOpen = true
		}
		for _, str := range monitorStrings {
			if strings.Contains(line, str.substr) {
				str.action(t, line)
				break
			}
		}
		// add to output buffer
		t.message.OutputBuf[t.message.outputPos] = line
		t.message.outputPos = (t.message.outputPos + 1) % len(t.message.OutputBuf)
		if t.message.outputSize < len(t.message.OutputBuf) {
			t.message.outputSize++
		}
		fmt.Println(line)
	}
}

func (t *TModProc) Kill() error {
	return t.proc.Process.Kill()
}

func CommandHandler(t *TModProc, cmd string) string {
	if t.IsOpen == false {
		return "Server is not running"
	}
	return HandleCmd(t, cmd)
}
