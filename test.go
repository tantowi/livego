//+build ignore

package main

import (
	"github.com/golang/glog"
	"os/exec"
)

func Bash(cmdline string) {
	cmdObj := exec.Command("bash", "-c", cmdline)
	output, err := cmdObj.CombinedOutput()
	if err != nil {
		if ins, ok := err.(*exec.ExitError); ok {
			out := string(output)
			exitcode := ins.ExitCode()
			glog.Warningf("bash cmd:[%s], out:[%s], exitcode:[%d]", cmdline, out, exitcode)
			return
		}
	}
	return
}

func main() {
	Bash("bash bootstrap.sh")
}