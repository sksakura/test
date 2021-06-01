package main

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*",
		strings.Repeat(" ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("*", j))
}

func doSomeWork(in int) {
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}
func main_1() {
	for i := 0; i < goroutinesNum; i++ {
		go doSomeWork(i)

	}
	fmt.Scanln()
}
