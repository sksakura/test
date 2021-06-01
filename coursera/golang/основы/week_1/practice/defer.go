package main

import (
	"fmt"
)

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened first: ", err)
			//panic(2)
		}
	}()
	fmt.Println("imitate work")
	panic("something happen")
}

func main34() {
	defer fmt.Println("DEFER function call 2")
	defer fmt.Println("DEFER function call 1")
	fmt.Println("some work")

	deferTest()
}
