package client

import (
	"main/options"
	"main/pkg/logger"
	"net"
	"time"
)

type pClient struct {
	conn   net.Conn
	status byte
	step   byte
}

// Client .
type Client interface {
	Ready()  // 5
	Finish() // 20
	Reset()  // 40
	Close()
}

func (p *pClient) heartBeat() {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			p.conn.Write([]byte{0, p.status, 0, p.step})
			p.setStatus()
		}
	}
}

func (p *pClient) setStatus() {
	if p.status == 1 {
		p.status = 0
	} else {
		p.status = 1
	}
}

func (p *pClient) Ready() {
	p.step = 5
}

func (p *pClient) Finish() {
	p.step = 14
}

func (p *pClient) Reset() {
	p.step = 28
	p.conn.Write([]byte{0, p.status, 0, p.step})
}

func (p *pClient) Close() {
	p.conn.Close()
}

// NewClient .
func NewClient(option *options.Option) Client {
	if option.Client == "" {
		option.Client = "192.168.0.10:2000"
	}
	cli, err := net.Dial("tcp", option.Client)
	if err != nil {
		logger.Printf("err is %v", err)
		panic(err)
	}

	p := &pClient{conn: cli}
	go p.heartBeat()
	return p
}
