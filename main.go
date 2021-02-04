package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"main/options"
	"main/server"

	"github.com/flopp/go-findfont"
)

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		// fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") || strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

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

	s := server.NewServer(o)
	s.Run()
}
