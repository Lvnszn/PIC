package logger

import (
	"fmt"
	"log"
)

// Printf .
func Printf(format string, v ...interface{}) {
	// 删除到这里
	fmt.Printf(format, v...)
	log.Printf(format, v...)
}
