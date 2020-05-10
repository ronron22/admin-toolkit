package main

import (
	"fmt"
	"github.com/prometheus/procfs"
	"log"
	"regexp"
	"time"
)

const (
	match_pattern = "php"
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

	var vms []uint
	var cpus []float64
	var starttimes []float64

	for _, a := range c {
		//fmt.Println(procfs.Proc(a.PID))
		//fmt.Println(a)
		stats, _ := a.NewStat()
		//fmt.Println(stats.Comm)
		aa := stats.Comm
		ab := stats.Starttime
		ac := stats.NumThreads
		amem := stats.VSize
		match, _ := regexp.MatchString(match_pattern, aa)
		if match {
			fmt.Println(aa)
			fmt.Println("starttime", ab)
			fmt.Println("thread", ac)
			fmt.Println("mem", amem)
			fmt.Printf("cpu time: %fs\n", stat.CPUTime())
			cpus = append(cpus, stat.CPUTime())
			fmt.Printf("vsize:    %dB\n", stat.VirtualMemory())
			vmem := stat.VirtualMemory()
			fmt.Println("vmem int64", int64(vmem))
			vms = append(vms, stat.VirtualMemory())
			fmt.Printf("rss:      %dB\n", stat.ResidentMemory())
			//fmt.Printf("%f", stat.StartTime())
			tt, _ := stat.StartTime()
			fmt.Printf("%f \n", tt)
			now := time.Now()
			secs := now.Unix()
			fmt.Println("Now", secs)
			fmt.Println("Start since", secs-int64(tt))

		}
	}
	fmt.Println(vms, cpus, starttimes)

	var total_vms uint = 0
	for _, vm := range vms {
		total_vms = total_vms + vm
	}
	fmt.Println("total memory", total_vms)
	fmt.Println("total memory", ((total_vms / 1024) / 1024))

}
