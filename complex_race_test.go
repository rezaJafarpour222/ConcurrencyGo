package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Complex(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	complex_race()
	_ = w.Close()
	result,_ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if(!strings.Contains(output, "$34320.00")){
		t.Error("wrong balance return")
	}
}
