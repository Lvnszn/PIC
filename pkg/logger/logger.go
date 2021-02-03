package logger

import (
	"log"
	"os"
)

// Printf .
func Printf(format string, v ...interface{}) {
	// 调试完成之后 删除这边的log
	f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	st, _ := f.Stat()
	f.Truncate(st.Size())
	log.SetOutput(f)
	// 删除到这里

	log.Printf(format, v...)
}
