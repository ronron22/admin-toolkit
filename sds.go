package main

import (
	"fmt"
	"os"
	//"syscall"
)

func IsRunning(i int) {
	proc, _ := os.FindProcess(i)

	fmt.Println(proc)
}

func main() {
	IsRunning(32403)
}
