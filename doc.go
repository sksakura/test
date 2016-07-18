// sks1 project doc.go

/*
sks1 document
*/
package sks1

import (
	"fmt"
)

func getFibvalue(n int) {
	fnc := fibbGenerator()
	for i := 0; i < n; i++ {
		fib := fnc()
		if i == n-1 {
			fmt.Println(fib)
		}
	}
}

func fibbGenerator() func() uint {
	i1 := uint(0)
	i2 := uint(1)
	return func() (k uint) {
		k = i1 + i2
		i1 = i2
		i2 = k
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func sum(args ...int) (sum int) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func half(i int) bool {
	return (i/2)%2 == 0
}
func max(args ...int) int {
	var max int = args[0]
	for i := 1; i < len(args); i++ {
		if max < args[i] {
			max = args[i]
		}
	}
	return max
}
