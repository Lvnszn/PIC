package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"main/options"
	"main/pkg/logger"
	"main/server"
)

func init() {
	os.Setenv("FYNE_FONT", "simkai.ttf")
}

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("配置信息")

	serverConfig := widget.NewEntry()
	serverConfig.SetPlaceHolder("192.168.0.10:2000")

	uri := widget.NewEntry()
	uri.SetPlaceHolder("username")

	port := widget.NewEntry()
	port.SetPlaceHolder(":2001")

	serverBox := widget.NewHBox(widget.NewLabel("PLC服务地址:"), layout.NewSpacer(), serverConfig)
	URIBox := widget.NewHBox(widget.NewLabel("数据库资源定位符:"), layout.NewSpacer(), uri)
	portBox := widget.NewHBox(widget.NewLabel("上位机端口:"), layout.NewSpacer(), port)

	startBtn := widget.NewButton("Run", func() {
		o := &options.Option{
			Client:   serverConfig.Text,
			Username: uri.Text,
			Server:   port.Text,
		}

		if o.Client == "" {
			o.Client = "192.168.0.10:2000"
		}

		fmt.Println("server endpoint:", serverConfig.Text, "uri:", uri.Text, "port: ", port.Text, "start")
		s := server.NewServer(o)
		go s.Run()
	})

	content := widget.NewVBox(serverBox, URIBox, portBox, startBtn)
	myWin.SetContent(content)
	logger.Printf("start show and run")
	myWin.ShowAndRun()
}
