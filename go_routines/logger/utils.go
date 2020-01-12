package logger

import (
	"bytes"
	"runtime"
	"strconv"
)

var (
	Debug   = false
	Verbose = false
)

func GetRoutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}
