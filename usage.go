package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)


func main() {
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{ nil, os.Stdout, os.Stderr }
	cmdh, err := os.StartProcess(os.Args[1], os.Args[1:], &procAttr)

	if err != nil {
		log.Fatal(err)
		return
	}

	status, err := cmdh.Wait()
	fmt.Println(status.SystemTime(), "sys")
	fmt.Println(status.UserTime(), "user")

	usage := status.SysUsage()
	fmt.Println(usage.(*syscall.Rusage).Maxrss, "kb RSS")
	fmt.Println(usage.(*syscall.Rusage).Nswap, "swaps")
	fmt.Println(usage.(*syscall.Rusage).Minflt, "minor faults")
	fmt.Println(usage.(*syscall.Rusage).Majflt, "major faults")
	fmt.Println(usage.(*syscall.Rusage).Nsignals, "signals delivered")
}

