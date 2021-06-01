package main

import "fmt"

func example1() {
	ch1 := make(chan int, 1)
	go func(in chan int) {
		val := <-in
		fmt.Println("GO: get from chan", val)
		fmt.Println("GO: after read from chan")
	}(ch1)
	ch1 <- 42
	fmt.Println("MAIN: after put to chan")

}

func example2() {
	in := make(chan int, 0)
	go func(out chan<- int) {
		for i := 0; i <= 10; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		close(out)
		fmt.Println("generator finish")
	}(in)
	for i := range in {
		fmt.Println("\tget", i)
	}
}
func example3() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)
	ch1 <- 1
	select {
	case val := <-ch1:
		fmt.Println("ch1 val", val)
	case ch2 <- 1:
		fmt.Println("put val to ch2")
	default:
		fmt.Println("default case")
	}
}

func example4() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch2 := make(chan int, 2)
	ch2 <- 3
LOOP:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("chan1 val", v1)
		case v2 := <-ch2:
			fmt.Println("chan2 val", v2)
		default:
			break LOOP
		}
	}
}

func example5() {
	cancelCh := make(chan struct{})
	dataCh := make(chan int)
	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh:
				return
			case dataCh <- val:
				val++
			}
		}
	}(cancelCh, dataCh)
	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- struct{}{}
			break
		}
	}
}

func main() {
	example5()
	fmt.Scanln()
}
