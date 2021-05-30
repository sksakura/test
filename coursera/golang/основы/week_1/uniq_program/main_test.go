package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk string = `1
2
3
4
5`
var testOkresult string = `1
2
3
4
5
`

var testFail string = `1
2
1`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("TestOk failed")
	}
	if out.String() != testOkresult {
		t.Errorf("TestOk failed. Results not mached:\n %v %v ", out.String(), testOkresult)

	}
}
func TestCheckErr(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("TestCheckErr failed")
	}
}
