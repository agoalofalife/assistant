package assistant

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
	"strconv"
)

type Processing interface {
	allPs() bool
}

type Convert interface {
	toJson() string
}
type containerProcess struct {
	listCurrentProcesses []Process
}
type Process struct {
	UID   int    `json:"UID"`
	PID   int    `json:"PID"`
	PPID  int    `json:"PPID"`
	F     int    `json:"F"`
	CPU   string `json:"CPU"`
	USER  string `json:"USER"`
	NI    int    `json:"NI"`
	RSS   int    `json:"RSS"`
	WCHAN string `json:"WCHAN"`
	STAT  string `json:"S"`
	STIME string `json:"STIME"`
	TIME  string `json:"TIME"`
	CMD   string `json:"CMD"`
	TTY   string `json:"TTY"`
	containerProcess
}

func New() *Process {
	return &Process{}
}

// list all process
func (process *Process) allPs() []Process {
	cmd, _ := exec.Command("ps", "o", "uid,pid,ppid,cpu,ni,f,stime,tty,rss,wchan,time,stat,user,command", "-A").Output()
	mapProcesses := make(map[int]map[string]interface{})
	middleProcess := make(map[string]interface{})

	// split all rows
	rows := strings.Split(string(cmd), "\n")

	//processes:= make([]Process, len(rows) -1 )
	titleString := strings.Fields(rows[0])

	// the idea is that the string COMMAND can have gaps > 0 and
	// and we leave it to the end

	for indexRow, rowValue := range rows[1:len(rows) - 1] {
		for index, attributeProcess := range strings.Fields(string(rowValue)) {
			if index >= len(titleString) {
				middleProcess[titleString[len(titleString)-1]] = middleProcess[titleString[len(titleString)-1]].(string) + " " + attributeProcess
			} else {
				middleProcess[titleString[index]] = attributeProcess
			}
		}
		mapProcesses[indexRow] = middleProcess
		middleProcess = make(map[string]interface{})
	}

	processList := make([]Process, len(mapProcesses) )

	for i := 0; i <= len(mapProcesses) - 1; i++ {
		pid, _  := strconv.Atoi(strings.TrimSpace(mapProcesses[i]["PID"].(string)))
		uid,_   := strconv.Atoi(mapProcesses[i]["UID"].(string))
		ppid,_  := strconv.Atoi(mapProcesses[i]["PPID"].(string))
		f,_     := strconv.Atoi(mapProcesses[i]["F"].(string))
		ni,_    := strconv.Atoi(mapProcesses[i]["NI"].(string))
		rss,_   := strconv.Atoi(mapProcesses[i]["RSS"].(string))
		var tty string
		if mapProcesses[i]["TTY"] == nil {
			tty = mapProcesses[i]["TT"].(string)
		} else{
			tty = mapProcesses[i]["TTY"].(string)
		}

		processList[i] = Process{PID: pid, TIME: mapProcesses[i]["TIME"].(string),
			CMD: mapProcesses[i]["COMMAND"].(string), TTY: tty,
			UID:uid,PPID:ppid, F:f, CPU:mapProcesses[i]["CPU"].(string),USER:mapProcesses[i]["USER"].(string),
			NI:ni, RSS:rss, WCHAN:mapProcesses[i]["WCHAN"].(string), STAT:mapProcesses[i]["STAT"].(string),
			STIME:mapProcesses[i]["STIME"].(string)}
	}

	process.containerProcess.listCurrentProcesses = processList
	return processList
}

func (container containerProcess) toJson() []byte {
	by, err := json.Marshal(container.listCurrentProcesses)
	if err != nil {
		log.Println(err)
	}
	return by
}
