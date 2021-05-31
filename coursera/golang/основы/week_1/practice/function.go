package main

import (
	"fmt"
)

func mainf() {
	printer := func(in string) {
		fmt.Println("print: ", in)
	}
	printer("hello world")

	type strfunctype func(string)
	worker := func(callback strfunctype) {
		callback("call callback")
	}
	worker(printer)

	//функция возвращает замыкание
	prefixer := func(prefix string) strfunctype {
		return func(in string) {
			fmt.Println(prefix, in)
		}
	}
	logger := prefixer("PREFIX ")
	logger("TEXT")
}
