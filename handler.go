package main

import (
	"github.com/polevpn/anyvalue"
	"github.com/polevpn/elog"
	"sync"
)

type RequestHandler struct {
	conn  Conn
	mutex *sync.Mutex
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{mutex: &sync.Mutex{}}
}

func (rh *RequestHandler) onCallback(av *anyvalue.AnyValue) {
	pkt, _ := av.EncodeJson()
	rh.conn.Send(pkt)
}

func (rh *RequestHandler) OnRequest(pkt []byte, conn Conn) {

	defer handlePanic()

	req, err := anyvalue.NewFromJson(pkt)

	if err != nil {
		elog.Error("decode json fail,", err)
		return
	}
	event := req.Get("event").AsStr()
	elog.Info("event=", event, ",req=", req)
	switch event {
	case "notification":
		data := req.Get("data")
		if data != nil {
			rh.onCallback(data)
		}
	}
}

func (rh *RequestHandler) OnConnected(conn Conn) {
	if rh.conn != nil {
		rh.conn.Close(true)
	}
	rh.conn = conn
}
