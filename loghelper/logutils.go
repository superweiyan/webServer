package loghelper

import (
	"io"
	"log"
	"os"
)

// it is for log
var (
	Info *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open error log file:", err)
	} else {
		log.Println("logutil init")
	}

	Info = log.New(io.MultiWriter(file, os.Stderr), "INFO:", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

//P is print
// func P(s string) {
// 	info.Println(s)
// }
