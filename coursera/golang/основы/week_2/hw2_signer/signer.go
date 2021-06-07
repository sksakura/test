package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex

func runJb(jb job, in chan interface{}, out chan interface{}) {
	//mu.Lock()
	jb(in, out)
	//mu.Unlock()
}

// сюда писать код
// Написание функции ExecutePipeline которая обеспечивает нам конвейерную обработку функций-воркеров
func ExecutePipeline(freeFlowJobs ...job) {
	mu = sync.Mutex{}
	var chans [100]chan interface{}
	openChans := 0

	for i, jb := range freeFlowJobs {
		chans[i] = make(chan interface{}, 100)
		openChans++

		var in chan interface{}
		if i > 0 {
			in = chans[i-1]
		}

		go runJb(jb, in, chans[i])

	}

	for {
		if len(chans[openChans-2]) > 0 {
			break
		}
	}

	for {
		readyChans := 0
		for i := 0; i < openChans; i++ {
			if len(chans[i]) == 0 {
				readyChans++
			} else {
				readyChans--
			}
			time.Sleep(time.Millisecond * 10)
		}
		if readyChans == openChans {
			break
		}
	}
	fmt.Println("KILL ALL CHANS")
	for i := 0; i < openChans; i++ {
		mu.Lock()
		close(chans[i])
		mu.Unlock()
	}
}

//считает значение crc32(data)+"~"+crc32(md5(data)) ( конкатенация двух строк через ~),
//где data - то что пришло на вход (по сути - числа из первой функции)
func SingleHash(in, out chan interface{}) {
	for val := range in {
		data := fmt.Sprint(val)
		md5 := DataSignerMd5(data)
		cr32md5 := DataSignerCrc32(md5)
		cr32 := DataSignerCrc32(data)
		result := cr32 + "~" + cr32md5
		out <- result
	}
}

//считает значение crc32(th+data)) (конкатенация цифры, приведённой к строке и строки),
//где th=0..5 ( т.е. 6 хешей на каждое входящее значение ), потом берёт конкатенацию результатов в порядке расчета (0..5),
// где data - то что пришло на вход (и ушло на выход из SingleHash)
func MultiHash(in, out chan interface{}) {
	for val := range in {
		data := fmt.Sprint(val)
		result := ""
		for i := 0; i < 6; i++ {
			th_data := strconv.Itoa(i) + data
			crc32 := DataSignerCrc32(th_data)
			result += crc32
		}
		out <- result
	}
}
func CombineResults(in, out chan interface{}) {

	for val := range in {
		fmt.Println(val)
	}
	fmt.Println("Now sort it")
}
