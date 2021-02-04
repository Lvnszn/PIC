package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"main/options"
	"main/server"
)

func main() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	o := &options.Option{}
	json.Unmarshal(b, &o)

	if o.Client == "" {
		o.Client = "192.168.0.10:2000"
	}

	log.Printf("%v", o)
	s := server.NewServer(o)
	go s.Run()

	g := make(chan os.Signal)
	signal.Notify(g, syscall.SIGTERM, syscall.SIGINT)
	sig := <-g
	log.Printf("caught sig: %+v, process will exit 2 seconds later..", sig)
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
