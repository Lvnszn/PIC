package server

import (
	"io"
	"main/pkg/database"
	"main/pkg/logger"
	"net"
	"strings"
)

type mseHandler struct {
	listener net.Listener
	db       database.DBClient
}

func (m *mseHandler) process() {
	for {
		conn, err := m.listener.Accept()
		if err != nil {
			logger.Printf("accept fail %v \n", err)
			continue
		}
		go func() {
			defer func() {
				conn.Close()
				err := recover()
				if err != nil {
					logger.Printf("err is %v", err)
				}
			}()

			for {
				buf := make([]byte, 256)
				n, err := conn.Read(buf[:])
				if err != nil {
					if err == io.EOF {
						continue
					}
					logger.Printf("read fail %v \n", err)
					return
				}

				req := string(buf[:n])
				logger.Printf("recive data is %s", req)
				if strings.Contains(req, "barcode#") {
					v := strings.Replace(req, "barcode#", "", -1)
					cnt, err := m.db.Select(v)
					if err != nil {
						logger.Printf("query from db err %v", err)
						return
					}

					logger.Printf("result is %v", cnt)
					if cnt == -1 {
						_, err = conn.Write([]byte("permit#NO"))
					} else if cnt == 1 {
						_, err = conn.Write([]byte("permit#OK"))
					} else {
						_, err = conn.Write([]byte("permit#NG"))
					}
					if err != nil {
						logger.Printf("write back err %v", err)
					}
				}
			}
		}()
	}
}

func NewMSEServe(addr string, db database.DBClient) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Printf("listen fail %v", err)
		return
	}

	handler := &mseHandler{}
	handler.listener = listener
	handler.db = db
	go handler.process()
}
