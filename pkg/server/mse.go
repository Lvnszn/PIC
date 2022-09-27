package server

import (
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
			buf := make([]byte, 256)
			_, err = conn.Read(buf[:])
			if err != nil {
				logger.Printf("read fail %v \n", err)
				return
			}

			req := string(buf)
			if strings.Contains(req, "barcode#") {
				v := strings.Replace(req, "barcode#", "", -1)
				cnt, err := m.db.Select(v)
				if err != nil {
					return
				}
				if cnt == 0 {
					conn.Write([]byte("permit#NO"))
					return
				} else if cnt > 0 {
					conn.Write([]byte("permit#OK"))
					return
				} else {
					conn.Write([]byte("permit#NG"))
					return
				}
			}
		}()
	}
}

func NewMSEServe(addr string, db database.DBClient) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	handler := &mseHandler{}
	handler.listener = listener
	handler.db = db
	go handler.process()
}
