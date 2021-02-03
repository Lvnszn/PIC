package logger

import (
	"log"
	"os"
)

// Printf .
func Printf(format string, v ...interface{}) {
	f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Printf(format, v...)
}
