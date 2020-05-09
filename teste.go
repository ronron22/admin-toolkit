package main

import (
	"fmt"
	"github.com/prometheus/procfs"
	"log"
	"regexp"
)

func main() {
	p, err := procfs.Self()
	if err != nil {
		log.Fatalf("could not get process: %s", err)
	}

	stat, err := p.NewStat()
	if err != nil {
		log.Fatalf("could not get process stat: %s", err)
	}

	fmt.Printf("command:  %s\n", stat.Comm)
	fmt.Printf("cpu time: %fs\n", stat.CPUTime())
	fmt.Printf("vsize:    %dB\n", stat.VirtualMemory())
	fmt.Printf("rss:      %dB\n", stat.ResidentMemory())

	c, _ := procfs.AllProcs()

	var vm []uint
	var cpu []float64
	var starttime []float64

	for _, a := range c {
		//fmt.Println(procfs.Proc(a.PID))
		//fmt.Println(a)
		stats, _ := a.NewStat()
		//fmt.Println(stats.Comm)
		aa := stats.Comm
		match, _ := regexp.MatchString("fpm", aa)
		if match {
			fmt.Println(aa)
			fmt.Printf("cpu time: %fs\n", stat.CPUTime())
			cpu = append(cpu, stat.CPUTime())
			fmt.Printf("vsize:    %dB\n", stat.VirtualMemory())
			vm = append(vm, stat.VirtualMemory())
			fmt.Printf("rss:      %dB\n", stat.ResidentMemory())
			//fmt.Printf("%f", stat.StartTime())
			tt, _ := stat.StartTime()
			fmt.Printf("%f \n", tt)
			starttime = append(starttime, tt)

		}
	}
	fmt.Println(vm, cpu, starttime)

}
