package main

type ProcessHandler interface {
	OnConnected(conn Conn)
	OnRequest(pkg []byte, conn Conn)
}
