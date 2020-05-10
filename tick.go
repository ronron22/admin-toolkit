package main

import "fmt"
import "time"

func main() {
	fmt.Println(time.Tick(1 * time.Second))
}
