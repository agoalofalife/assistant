package assistant

import (
	"os/exec"
	"regexp"
	"github.com/assistant/helpers"
	"strconv"
	"strings"
	"encoding/json"
	"log"

	"os"
)


type Processing interface {
	allPs() bool
}

type Convert interface {
	toJson() string
}
type containerProcess struct{
	listCurrentProcesses []Process
}
type Process struct {
	UUD int `json:"UID"`
	PID int `json:"PID"`
	PPID int `json:"PPID"`
	F int `json:"F"`
	CPU string `json:"CPU"`
	PRI int `json:"PRI"`
	NI int `json:"NI"`
	SZ int `json:"SZ"`
	RSS int `json:"RSS"`
	WCHAN string `json:"WCHAN"`
	S string `json:"S"`
	ADDR int `json:"ADDR"`
	STIME string `json:"STIME"`
	TIME string `json:"TIME"`
	CMD string `json:"CMD"`
	TTY string `json:"TTY"`
	containerProcess
}

func New() *Process {
	return &Process{}
}
// list all process
func (process *Process) allPs() []Process {
	cmd ,_:= exec.Command("ps","-lf").Output()

	titleString := strings.Split(string(cmd), "\n")[0]
	//log.Println(strings.Fields(titleString))
	for _,nameFields := range strings.Fields(titleString) {
		log.Println(nameFields)
	}
	os.Exit(2)
	regexTerminal := regexp.MustCompile("\\s[0-9]+\\s([A-z0-9]+|.+)\\s[0-9]+:[0-9]+\\.[0-9]+")
	regexPid := regexp.MustCompile("([0-9]+)\\s.+")
	regexTime := regexp.MustCompile("[0-9]+:[0-9]+\\.[0-9]+")
	regexPath := regexp.MustCompile("[0-9]+:[0-9]+\\.[0-9]+\\s+(.+)")

	listTerminal := helpers.GroupExclude(regexTerminal.FindAllStringSubmatch(string(cmd), -1))
	listPid := helpers.GroupExclude(regexPid.FindAllStringSubmatch(string(cmd), -1))
	listTimes := regexTime.FindAllString(string(cmd), -1)
	listPaths := helpers.GroupExclude(regexPath.FindAllStringSubmatch(string(cmd), -1))

	processList := make([]Process, len(listPid))
	for i := 0; i <= len(listPid) - 1; i++{
		pid ,_:= strconv.Atoi(strings.TrimSpace(listPid[i]))
		processList[i] = Process{PID:pid, TIME:listTimes[i], CMD:listPaths[i],TTY:listTerminal[i]}
	}

	process.containerProcess.listCurrentProcesses = processList
	return processList
}

func (container containerProcess)  toJson () ([]byte){
	by, err := json.Marshal(container.listCurrentProcesses)
	if err !=nil {
		log.Println(err)
	}
	return by
}