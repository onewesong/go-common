package os

import (
	"os"
	"syscall"
)

func IsPidRunning(pid int) bool {
	process, err := os.FindProcess(int(pid))
	if err != nil {
		return false
	} else {
		err := process.Signal(syscall.Signal(0))
		if err != nil {
			return false
		}
		return true
	}
}
