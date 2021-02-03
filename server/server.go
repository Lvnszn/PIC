package server

import (
	"encoding/hex"
	"main/client"
	"main/options"
	"main/pkg/database"
	"main/pkg/logger"
	"main/pkg/parser"
	"main/pkg/protocol"
	"net"
)

type pServer struct {
	cli   client.Client
	dbCli database.DBClient
	port  string
}

// Server .
type Server interface {
	Run()
}

// Run .
func (p *pServer) Run() {
	listener, err := net.Listen("tcp", p.port)
	if err != nil {
		logger.Printf("listen %v fail: %v", p.port, err)
		return
	}
	defer listener.Close()
	for {
		p.cli.Ready()
		connection, err := listener.Accept()
		if err != nil {
			logger.Printf("Accept 失败: %v", err)
			continue
		}
		p.process(connection)
	}
}

func (p *pServer) process(conn net.Conn) {
	defer conn.Close()
	for {
		var b [1024]byte
		n, err := conn.Read(b[:])
		if err != nil {
			logger.Printf("read fail: %v", err)
			break
		}
		hexStr := hex.EncodeToString(b[:])
		logger.Printf("hex string is: %s", hex.EncodeToString(b[:]))
		logger.Printf("bytes is: %v", b)

		start, end := parser.IdxOfHead(hexStr)
		if end > n {
			end = n
		}
		logger.Printf(hexStr[start+20 : start+22])

		if parser.IsProcess(hexStr[start+20 : start+22]) {
			logger.Printf("status is process and write to db %v", hexStr[start+20:start+22])
			entity := protocol.DecodeMsg(b[start:end])
			sql := entity.GenSQL()
			err := p.dbCli.Insert(sql)
			p.cli.Finish()
			if err != nil {
				logger.Printf("insert error %v", err)
				return
			}
		} else if parser.AckFinish(hexStr[start+20 : start+22]) {
			p.cli.Reset()
			return
		}
	}
}

// NewServer receive bytes from plc
func NewServer(opt *options.Option) Server {
	cli := client.NewClient(opt)
	if opt.Server == "" {
		opt.Server = "0.0.0.0:2001"
	}
	db := database.NewMssql(opt)
	p := &pServer{
		cli:   cli,
		dbCli: db,
		port:  opt.Server,
	}
	return p
}
