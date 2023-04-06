package process

import (
	"log"
	"time"
)

const (
	SERVER_STARTED  = "Server started"
	HAS_JOINED      = "has joined"
	HAS_LEFT        = "has left"
	HELP            = "help"
	PLAYING         = "playing"
	CLEAR           = "clear"
	EXIT            = "exit"
	EXIT_NOSAVE     = "exit-nosave"
	SAVE            = "save"
	KICK            = "kick"
	BAN             = "ban"
	PASSWORD_SHOW   = "password"
	PASSWORD_CHANGE = "password set"
	VERSION         = "version"
	TIME            = "time"
	PORT            = "port"
	MAXPLAYERS      = "maxplayers"
	SAY             = "say"
	MOTD_SHOW       = "motd"
	MOTD_CHANGE     = "motd set"
	DAWN            = "dawn"
	NOON            = "noon"
	DUSK            = "dusk"
	MIDNIGHT        = "midnight"
	SETTLE          = "settle"
	SEED            = "seed"
	MODLIST         = "modlist"
)

var (
	// no return value required
	notReturnMap map[string]func(*TModProc, string) string
	// not used
	notUsedMap map[string]func(*TModProc, string) string
	// require return value
	requireReturnMap map[string]func(*TModProc, string) string
	// unknown command
	unknownCmd = func(t *TModProc, cmd string) string {
		return "unknown command"
	}
)

func init() {
	noReturn := func(t *TModProc, cmd string) string {
		if _, err := t.inPipe.Write([]byte(cmd)); err != nil {
			log.Printf("Error writing to pipe: %s......", err.Error())
		}
		return "command complete"
	}
	notReturnMap = map[string]func(*TModProc, string) string{
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
	notUsed := func(t *TModProc, cmd string) string {
		return "command not used"
	}
	notUsedMap = map[string]func(*TModProc, string) string{
		HELP:       notUsed,
		PLAYING:    notUsed,
		CLEAR:      notUsed,
		PORT:       notUsed,
		MAXPLAYERS: notUsed,
	}
	requireReturn := func(t *TModProc, cmd string) string {
		if _, err := t.inPipe.Write([]byte(cmd)); err != nil {
			log.Printf("Error writing to pipe: %s......", err.Error())
		}
		time.Sleep(500 * time.Millisecond) // wait for the command to produce output
		out := t.message.OutputBuf
		lastLine := out[len(out)-1]
		return lastLine
	}
	requireReturnMap = map[string]func(*TModProc, string) string{
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
func HandleCmd(t *TModProc, cmd string) string {
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
		t.ps.players[line] = now
		t.ps.count++
	}},
	{HAS_LEFT, func(t *TModProc, line string) {
		delete(t.ps.players, line)
		t.ps.count--
	}},
}
