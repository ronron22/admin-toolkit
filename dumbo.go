package main

import (
	"fmt"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/prometheus/procfs"
	"log"
	"os"
	"regexp"
	"time"
)

const (
	match_pattern = "php"
)

func main() {

	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}

	c, err := procfs.AllProcs()
	if err != nil {
		log.Fatalf("could not listed process: %s", err)
		os.Exit(2)
	}

	var vms []uint
	var swappeds []uint64

	var Start_time int64
	counter := 0
	for _, a := range c {
		stats, err := a.Stat()
		if err != nil {
			log.Fatalf("could not get process stat: %s", err)
			os.Exit(2)
		}

		status, _ := a.NewStatus()
		if err != nil {
			log.Fatalf("could not get process status: %s", err)
			os.Exit(2)
		}
		aa := stats.Comm
		match, _ := regexp.MatchString(match_pattern, aa)
		if match {
			counter += 1
			fmt.Printf("Name:        %s, PID: %d\n", aa, stats.PID)
			fmt.Printf("cpu time:    %fs\n", stats.CPUTime())
			fmt.Printf("vsize:       %dB\n", (stats.VirtualMemory() / 1024))
			fmt.Printf("rss:         %dB\n", (stats.ResidentMemory() / 1024))
			swapped := status.VmSwap
			fmt.Printf("swapped      %dB\n", swapped)
			swappeds = append(swappeds, swapped)
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
	var total_swappeds uint64 = 0
	for _, s := range swappeds {
		total_swappeds = total_swappeds + s
	}
	fmt.Printf("--- digest ---\n")
	fmt.Printf("memory total: %d Mo\n", ((memory.Total / 1024) / 1034))
	fmt.Printf("memory free: %d Mo\n", ((memory.Free / 1024) / 1024))
	fmt.Printf("[1] there %d processus %s \n", counter, match_pattern)
	fmt.Printf("[2] with a total memory of %d Mo (rss)\n", (total_vms / 1024 / 1024))
	fmt.Printf("[3] with a total of %d Ko swapped\n", (total_swappeds / 1024))
	switch {
	case Start_time > 86400:
		fmt.Printf("[4] started since %d d \n", ((Start_time/60)/60)/60)
	case Start_time > 3600:
		fmt.Printf("[4] started since %d h \n", (Start_time/60)/60)
	case Start_time > 60:
		fmt.Printf("[4] started since %d m \n", Start_time/60)
	default:
		fmt.Printf("[4] started since %d s \n", Start_time)
	}
	fmt.Printf("--- end ---\n")

}
