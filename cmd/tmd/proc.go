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
	ID       int
	proc     *exec.Cmd
	inPipe   io.WriteCloser
	outPipe  io.ReadCloser
	message  *OutputMessage
	ps       *PlayerStatus
	stopChan chan bool
	IsOpen   bool
}

func NewTModLoader(config string, id int) (*TModProc, error) {
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

	t := &TModProc{
		ID:       id,
		proc:     proc,
		inPipe:   stdin,
		outPipe:  stdout,
		stopChan: make(chan bool),
		IsOpen:   false,
		message: &OutputMessage{
			OutputBuf:  [100]string{},
			outputPos:  0,
			outputSize: 0,
		},
		ps: &PlayerStatus{
			players: make(map[string]time.Time),
			count:   0,
		},
	}
	// add the new process to the list
	return t, nil
}

func (t *TModProc) Start() error {

	if err := t.proc.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		return err
	}

	// send a newline to the process to start it
	_, _ = t.inPipe.Write([]byte("\n"))

	go ServerMonitor(t)

	t.Wait()
	return nil
}

func (t *TModProc) Wait() error {
	// wait for the process to exit
	if err := t.proc.Wait(); err != nil {
		log.Printf("Error waiting for command: %s......", err.Error())
		return err
	}

	// close the pipes
	if err := t.inPipe.Close(); err != nil {
		log.Printf("Error closing inPipe: %s......", err.Error())
		return err
	}
	if err := t.outPipe.Close(); err != nil {
		log.Printf("Error closing outPipe: %s......", err.Error())
		return err
	}
	return nil
}

func (t *TModProc) Stop() error {
	// send the exit command to the process
	CommandHandler(t, EXIT)
	// wait for the process to exit
	if err := t.proc.Wait(); err != nil {
		log.Printf("Error waiting for command: %s......", err.Error())
		return err
	}
	// save the context to the process
	return t.proc.Process.Kill()
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

func CommandHandler(t *TModProc, cmd string) string {
	if t.IsOpen == false {
		return "Server is not running"
	}
	return HandleCmd(t, cmd)
}
