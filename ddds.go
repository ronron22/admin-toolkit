package main

import (
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"log"
	"os/exec"
)

func main() {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	// map ages
	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())

		// do os.* stuff on the pid
	}

	out := exec.Command("ps aux")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

}
