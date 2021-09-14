package logger

import (
	"log"
)

// Printf .
func Printf(format string, v ...interface{}) {
	// 删除到这里
	log.Printf(format, v...)
	log.Println()
}
