package assistant

//import "github.com/shirou/gopsutil"
import (
	"os/exec"
	"log"
	"os"
	//process "github.com/mitchellh/go-ps"
	"regexp"
	"fmt"
	"github.com/assistant/helpers"
)


type Process interface {
	allPs() bool
}

// list all process
func allPs() bool {
	//processes,_ := process.Processes()
	//for _,ps := range processes{
	//
	//	log.Println(os.FindProcess(ps.Pid()))
	//	os.Exit(2)
	//}
	cmd ,_:= exec.Command("ps").Output()
	regexTerminal := regexp.MustCompile("\\s[0-9]+\\s([A-z0-9]+)")

	log.Println(helpers.GroupExclude(regexTerminal.FindAllStringSubmatch(string(cmd), -1)))
	//regexTerminal.FindAllStringSubmatch(string(cmd), -1)
	//log.Println(regexTerminal.SubexpNames())
	os.Exit(2)
	regexPid := regexp.MustCompile("\\s[0-9]+\\s")
	listPid := regexPid.FindAllString(string(cmd), -1)

	fmt.Println( listPid )
	//log.Println(string(cmd))
	os.Exit(2)
	for st := range cmd{
		log.Println(string(st))

	}

	os.Exit(2)

	return true
}

