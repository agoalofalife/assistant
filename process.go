package assistant

import (
	"os/exec"
	"regexp"
	"github.com/assistant/helpers"
	"strconv"
	"strings"
)


type Processing interface {
	allPs() bool
}
type Process struct {
	PID int
	PPID int
	TIME string
	CMD string
	TTY string
}

// list all process
func allPs() []Process {
	cmd ,_:= exec.Command("ps").Output()

	regexTerminal := regexp.MustCompile("\\s[0-9]+\\s([A-z0-9]+)")
	regexPid := regexp.MustCompile("\\s[0-9]+\\s")
	regexTime := regexp.MustCompile("[0-9]{1,2}:[0-9]{1,2}\\.[0-9]{1,2}")
	regexPath := regexp.MustCompile("[0-9]{1,2}:[0-9]{1,2}\\.[0-9]{1,2}\\s+(.+)")

	listTerminal := helpers.GroupExclude(regexTerminal.FindAllStringSubmatch(string(cmd), -1))
	listPid := regexPid.FindAllString(string(cmd), -1)
	listTimes := regexTime.FindAllString(string(cmd), -1)
	listPaths := helpers.GroupExclude(regexPath.FindAllStringSubmatch(string(cmd), -1))

	processList := make([]Process, len(listPid) + 1)
	for i := 0; i <= len(listPid) -  1; i++{
		pid ,_:= strconv.Atoi(strings.TrimSpace(listPid[i]))
		processList[i] = Process{PID:pid, TIME:listTimes[i], CMD:listPaths[i],TTY:listTerminal[i]}
	}
	return processList
}
