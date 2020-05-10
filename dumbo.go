package main

import (
	"fmt"
	"github.com/prometheus/procfs"
	//"log"
	"regexp"
	"time"
)

const (
	match_pattern = "php"
)

func main() {

	c, _ := procfs.AllProcs()

	var vms []uint
	//var starttimes []float64

	var Start_time int64
	counter := 0
	for _, a := range c {
		//stats, _ := a.NewStat()
		stats, _ := a.Stat()
		status, _ := a.NewStatus()
		aa := stats.Comm
		//vmem := stats.VSize
		match, _ := regexp.MatchString(match_pattern, aa)
		if match {
			counter += 1
			fmt.Printf("Name:        %s, PID: %d\n", aa, stats.PID)
			fmt.Printf("cpu time:    %fs\n", stats.CPUTime())
			fmt.Printf("vsize:       %dB\n", (stats.VirtualMemory() / 1024))
			fmt.Printf("rss:         %dB\n", (stats.ResidentMemory() / 1024))
			fmt.Printf("swapped      %dB\n", status.VmSwap)

			vms = append(vms, uint(stats.ResidentMemory()))
			tt, _ := stats.StartTime()
			now := time.Now()
			secs := now.Unix()

			Start_time = secs - int64(tt)
			if Start_time > 60 {
				fmt.Printf("Start since: %d minute(s)\n\n", Start_time/60)
			}

		}
	}

	var total_vms uint = 0
	for _, vm := range vms {
		total_vms = total_vms + vm
	}
	fmt.Printf("--- digest ---\n")
	fmt.Printf("[1] there %d processus %s \n", counter, match_pattern)
	fmt.Printf("[2] with a total memory of %d Mo\n", (total_vms / 1024 / 1024))
	switch {
	case Start_time > 86400:
		fmt.Printf("[3] started since %d d \n", (Start_time / 60 / 60 / 60))
	case Start_time > 3600:
		fmt.Printf("[3] started since %d h \n", (Start_time / 60 / 60))
	case Start_time > 60:
		fmt.Printf("[3] started since %d m \n", Start_time/60)
	default:
		fmt.Printf("[3] started since %d s \n", Start_time)
	}
	fmt.Printf("--- end ---\n")

}
