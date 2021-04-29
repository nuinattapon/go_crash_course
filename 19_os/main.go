package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker run <container> cmd args
// go run main.go cmd args
func main() {

	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "run":
			run()

		default:
			panic("What??")
		}
	} else {
		panic("Pls specify a command to run")
	}
}

func run() {
	// cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	// }

	must(cmd.Run())

}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	// }
	must(cmd.Run())

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
