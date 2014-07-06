// +build linux darwin
// +build !cgo !windows

package safekill

// import "os"
import "syscall"

// send SIGUSR2 if on unix-like system
var HandleSignal = syscall.SIGUSR2

func SafeKill(pid int) error {
	// SIGINT is caught by both parent process and child process
	// but child process will handle it so that installation can go on
	return syscall.Kill(pid, HandleSignal)
}
