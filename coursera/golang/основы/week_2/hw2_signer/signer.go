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

func cr32md5Calc(out chan string, val string, wg *sync.WaitGroup) {
	mu.Lock()
	md5 := DataSignerMd5(val)
	mu.Unlock()
	result := DataSignerCrc32(md5)
	out <- result
	wg.Done()
	close(out)
	//fmt.Println("cr32md5Calc = ", result)
}

func Cr32Calc(out chan string, val string, wg *sync.WaitGroup) {
	result := DataSignerCrc32(val)
	out <- result
	close(out)
	wg.Done()
	//fmt.Println("Cr32Calc = ", result)
}

//считает значение crc32(data)+"~"+crc32(md5(data)) ( конкатенация двух строк через ~),
//где data - то что пришло на вход (по сути - числа из первой функции)
func SingleHash(in, sout chan interface{}) {
	wg := &sync.WaitGroup{}
	for val := range in {

		data := fmt.Sprint(val)
		wg.Add(1)
		go func(out chan interface{}, wg *sync.WaitGroup) {
			cr32md5Chan := make(chan string, 1)
			cr32Chan := make(chan string, 1)

			wgS := &sync.WaitGroup{}
			wgS.Add(1)
			go cr32md5Calc(cr32md5Chan, data, wgS)

			wgS.Add(1)
			go Cr32Calc(cr32Chan, data, wgS)

			wgS.Wait()

			cr32md5 := <-cr32md5Chan
			cr32 := <-cr32Chan

			result := cr32 + "~" + cr32md5
			//fmt.Println("result = ", result)
			out <- result
			wg.Done()
		}(sout, wg)
	}
	wg.Wait()

}

//считает значение crc32(th+data)) (конкатенация цифры, приведённой к строке и строки),
//где th=0..5 ( т.е. 6 хешей на каждое входящее значение ), потом берёт конкатенацию результатов в порядке расчета (0..5),
// где data - то что пришло на вход (и ушло на выход из SingleHash)
func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for val := range in {
		data := fmt.Sprint(val)
		wg.Add(1)
		go func(out chan interface{}, wg *sync.WaitGroup) {
			result := ""

			var chans [6]chan string
			wgS := &sync.WaitGroup{}
			for i := 0; i < 6; i++ {
				chans[i] = make(chan string, 1)
				wgS.Add(1)
				go Cr32Calc(chans[i], strconv.Itoa(i)+data, wgS)
			}
			wgS.Wait()

			for i := 0; i < 6; i++ {
				crc32 := <-chans[i]
				result += crc32
			}
			out <- result
			wg.Done()
		}(out, wg)
	}
	wg.Wait()
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
