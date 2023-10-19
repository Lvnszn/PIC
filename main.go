package main

import (
	"encoding/json"
	"io/ioutil"
	"main/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"main/client"
	"main/options"
	"main/pkg/database"
)

func init() {
	logger.Init("", "debug", "PIC")
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
	logger.Infof("%+v", o)

	db := database.NewMssql(o)
	client.NewClient(o, db)

	g := make(chan os.Signal)
	signal.Notify(g, syscall.SIGTERM, syscall.SIGINT)
	sig := <-g
	logger.Infof("caught sig: %+v, process will exit 2 seconds later..", sig)
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
