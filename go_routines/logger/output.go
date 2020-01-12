package logger

import (
	"fmt"
)

func Log(format string, params ...interface{}) {
	if !Debug {
		return
	}
	fmt.Printf(format, params...)
}
