package server

import (
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	NewMSEServe("0.0.0.0:8089", nil)
	c, er := net.Dial("tcp", "127.0.0.1:8089")
	if er != nil {
		panic(er)
	}

	c.Write([]byte("barcode#test"))
	buf := make([]byte, 100)
	//c.Read(buf)
	println(string(buf))
	c.Write([]byte("barcode#test"))
	time.Sleep(time.Second * 3)
}
