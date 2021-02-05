package client

import (
	"encoding/hex"
	"main/options"
	"main/pkg/database"
	"main/pkg/logger"
	"main/pkg/parser"
	"main/pkg/protocol"
	"net"
	"time"
)

type pClient struct {
	conn   net.Conn
	dbCli  database.DBClient
	status byte
	step   byte
}

// Client .
type Client interface {
	Close()
}

func (p *pClient) heartBeat() {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			p.conn.Write([]byte{p.status, 0, 0, p.step})
			p.setStatus()
			// p.process(p.conn)
		}
	}
}

func (p *pClient) consume() {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			p.process(p.conn)
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
	p.step = 20
}

func (p *pClient) IsFinish() bool {
	return p.step == 20
}

func (p *pClient) Reset() {
	p.step = 40
	p.conn.Write([]byte{p.status, 0, 0, p.step})
}

func (p *pClient) Close() {
	p.conn.Close()
}

func (p *pClient) process(conn net.Conn) {
	if !p.IsFinish() {
		p.Ready()
	}
	logger.Printf("ready times")
	var b [1024]byte
	n, err := conn.Read(b[:])
	if err != nil {
		logger.Printf("read fail: %v", err)
		return
	}
	hexStr := hex.EncodeToString(b[:])
	logger.Printf("hex string is: %s", hex.EncodeToString(b[:]))

	start, end := parser.IdxOfHead(hexStr)
	if end > n {
		end = n
	}
	logger.Printf("start index: %v", start)
	logger.Printf("hex status is: %s", hexStr[start+18:start+24])

	if !p.IsFinish() && parser.IsProcess(hexStr[start+18:start+20]) {
		logger.Printf("status is process and write to db %v", hexStr[start+18:start+20])
		entity := protocol.DecodeMsg(b[start:end])
		sql := entity.GenSQL()
		logger.Printf("insert sql is %v", sql)
		err := p.dbCli.Insert(sql)
		p.Finish()
		if err != nil {
			logger.Printf("insert error %v", err)
		}
	} else if p.IsFinish() && parser.AckFinish(hexStr[start+18:start+20]) {
		p.Reset()
	} else if parser.IsFine(hexStr[start+18 : start+20]) {
		p.Ready()
	}
}

// NewClient .
func NewClient(option *options.Option, db database.DBClient) Client {
	if option.Client == "" {
		option.Client = "192.168.0.10:2000"
	}
	cli, err := net.Dial("tcp", option.Client)
	if err != nil {
		logger.Printf("err is %v, server is not exists", err)
		panic(err)
	}

	p := &pClient{
		conn:  cli,
		dbCli: db,
	}
	go p.heartBeat()
	go p.consume()
	return p
}
