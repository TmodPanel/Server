package process

import (
	"errors"
	"log"
	"time"
)

var (
	// no return value required
	notReturnMap map[string]func(*TModProc, string) (string, error)
	// not used
	notUsedMap map[string]func(*TModProc, string) (string, error)
	// require return value
	requireReturnMap map[string]func(*TModProc, string) (string, error)
	// unknown command
	unknownCmd = func(t *TModProc, cmd string) (string, error) {
		return "", errors.New("unknown command")
	}
)

func init() {
	noReturn := func(t *TModProc, cmd string) (string, error) {
		if _, err := t.inPipe.Write([]byte(cmd)); err != nil {
			log.Printf("Error writing to pipe: %s......", err.Error())
			return "", err
		}
		return "command complete", nil
	}
	notReturnMap = map[string]func(*TModProc, string) (string, error){
		DAWN:        noReturn,
		NOON:        noReturn,
		DUSK:        noReturn,
		MIDNIGHT:    noReturn,
		EXIT:        noReturn,
		EXIT_NOSAVE: noReturn,
		SAVE:        noReturn,
		//KICK:        noReturn,
		//BAN:         noReturn,
	}
	notUsed := func(t *TModProc, cmd string) (string, error) {
		return "command not used", nil
	}
	notUsedMap = map[string]func(*TModProc, string) (string, error){
		HELP:       notUsed,
		PLAYING:    notUsed,
		CLEAR:      notUsed,
		PORT:       notUsed,
		MAXPLAYERS: notUsed,
	}
	requireReturn := func(t *TModProc, cmd string) (string, error) {
		if _, err := t.inPipe.Write([]byte(cmd)); err != nil {
			log.Printf("Error writing to pipe: %s......", err.Error())
			return "", err
		}
		time.Sleep(500 * time.Millisecond) // wait for the command to produce output
		out := t.message.OutputBuf
		lastLine := out[len(out)-1]
		return lastLine, nil
	}
	requireReturnMap = map[string]func(*TModProc, string) (string, error){
		VERSION:       requireReturn,
		TIME:          requireReturn,
		SAY:           requireReturn,
		MOTD_SHOW:     requireReturn,
		PASSWORD_SHOW: requireReturn,
		MODLIST:       requireReturn,
		//PASSWORD_CHANGE: requireReturn,
		MOTD_CHANGE: requireReturn,
		SEED:        requireReturn,
		//SETTLE:          requireReturn,
	}

}
func HandleCmd(t *TModProc, cmd string) (string, error) {
	// which class of command is it?
	if _, ok := notReturnMap[cmd]; ok {
		return notReturnMap[cmd](t, cmd)
	}
	if _, ok := notUsedMap[cmd]; ok {
		return notUsedMap[cmd](t, cmd)
	}
	if _, ok := requireReturnMap[cmd]; ok {
		return requireReturnMap[cmd](t, cmd)
	}
	return unknownCmd(t, cmd)
}

var monitorStrings = []struct {
	substr string
	action func(*TModProc, string)
}{
	{HAS_JOINED, func(t *TModProc, line string) {
		now := time.Now()
		ps := PlayerStatus{
			Name:     line,
			JoinTime: now,
		}
	}},
	{HAS_LEFT, func(t *TModProc, line string) {
		now := time.Now()

	}},
}
