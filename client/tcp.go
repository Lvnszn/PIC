package client

import (
	"encoding/hex"
	"io"
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

	date       string
	clientMeta string
}

// Client .
type Client interface {
	Close()
}

func (p *pClient) isContinue() bool {
	time.LoadLocation("Aisa/Shanghai")
	now := time.Now()
	logger.Printf("now date is %v", now)
	dd, _ := time.Parse("2006-01-02 15:04:05", p.date)
	logger.Printf("last date is %v", dd)
	return now.After(dd)
}

func (p *pClient) heartBeat() {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			if p.isContinue() {
				logger.Printf("%v is continue...", time.Now())
				continue
			}

			_, err := p.conn.Write([]byte{p.status, 0, 0, p.step})
			if err != nil {
				p.conn, _ = net.Dial("tcp", p.clientMeta)
				continue
			}
			p.setStatus()
		}
	}
}

func (p *pClient) consume() {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			if p.isContinue() {
				logger.Printf("%v is continue...", time.Now())
				continue
			}

			p.process()
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

func (p *pClient) StepToString() string {
	switch p.step {
	case 5:
		return "就绪"
	case 20:
		return "结束"
	}

	return ""
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

func (p *pClient) process() {
	if !p.IsFinish() {
		p.Ready()
	}
	logger.Printf("step is %v", p.StepToString())
	var b [1024]byte
	n, err := p.conn.Read(b[:])
	if err != nil {
		if err != io.EOF {
			p.conn, _ = net.Dial("tcp", p.clientMeta)
		}
		logger.Printf("read fail: %v", err)
		return
	}
	hexStr := hex.EncodeToString(b[:])

	start, end := parser.IdxOfHead(hexStr)
	if end > n {
		end = n
	}

	if !p.IsFinish() && parser.IsProcess(hexStr[start+18:start+20]) {
		logger.Printf("获取数据成功，正在处理...")
		entity := protocol.DecodeMsg(b[start:end])
		sql := entity.GenSQL()
		err := p.dbCli.Insert(sql)
		if err != nil {
			logger.Printf("insert error %v", err)
			return
		}
		p.Finish()
		logger.Printf("写入数据成功", p.StepToString())
	} else if p.IsFinish() && parser.AckFinish(hexStr[start+18:start+20]) {
		logger.Printf("接收到确认写入成功的信息，状态重置")
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
		conn:       cli,
		dbCli:      db,
		clientMeta: option.Client,
		date:       option.Date,
	}
	logger.Printf("开始发送心跳")
	go p.heartBeat()
	logger.Printf("开始接收PLC的数据")
	go p.consume()
	return p
}
