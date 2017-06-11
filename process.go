package assistant

//import "github.com/shirou/gopsutil"
import (
	"os/exec"
	"log"
	"os"
)
//import  "github.com/shirou/gopsutil"

type Process interface {
	allPs() bool
}

// list all process
func allPs() bool {
	cmd ,_:= exec.Command("ps").Output()
	log.Println(string(cmd[4]))
	os.Exit(2)
	for st := range cmd{
		log.Println(string(st))

	}

	os.Exit(2)

	return true
}

