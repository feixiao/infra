package net

import (
	"net"
	"runtime"
	"strings"

	lg "github.com/feixiao/infra/log"
)

// TCP长连接的处理只需要实现Handle即可
type TCPHandler interface {
	Handle(net.Conn)
}

// 首先是监听端口，当有请求到来时开启一个goroutine去处理该链接请求
func TCPServer(listener net.Listener, handler TCPHandler) {

	lg.Info("TCP: listening on %s", listener.Addr())

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				lg.Info("NOTICE: temporary Accept() failure - %s", err)
				runtime.Gosched()
				continue
			}
			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				lg.Info("ERROR: listener.Accept() - %s", err)
			}
			break
		}
		// TCPHandler接口调用
		go handler.Handle(clientConn)
	}

	lg.Info("TCP: closing %s", listener.Addr())
}
