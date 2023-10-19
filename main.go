package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"main/client"
	"main/options"
	"main/pkg/database"
)

func init() {
	file := "./" + "message" + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[PLC]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	o := &options.Option{}
	json.Unmarshal(b, &o)

	time.LoadLocation("Aisa/Shanghai")
	if o.Client == "" {
		o.Client = "192.168.0.10:2000"
	}
	log.Printf("%v", o)

	db := database.NewMssql(o)
	client.NewClient(o, db)

	g := make(chan os.Signal)
	signal.Notify(g, syscall.SIGTERM, syscall.SIGINT)
	sig := <-g
	log.Printf("caught sig: %+v, process will exit 2 seconds later..", sig)
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
