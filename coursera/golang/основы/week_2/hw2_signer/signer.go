package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex

func runJb(wg *sync.WaitGroup, jb job, in chan interface{}, out chan interface{}) {

	jb(in, out)
	wg.Done()
	close(out)
	//mu.Unlock()
}

// сюда писать код
// Написание функции ExecutePipeline которая обеспечивает нам конвейерную обработку функций-воркеров
func ExecutePipeline(freeFlowJobs ...job) {

	wg := &sync.WaitGroup{}
	defer wg.Wait()

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

		wg.Add(1)
		go runJb(wg, jb, in, chans[i])

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
	arr := make([]string, 0)
	for val := range in {
		data := fmt.Sprint(val)
		arr = append(arr, data)
	}
	sort.Strings(arr)
	result := strings.Join(arr[:], "_")

	out <- result

}
