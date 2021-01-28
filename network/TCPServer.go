package network

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/SERV4BIZ/gfp/handler"
)

// TCPServer is structure object
type TCPServer struct {
	Listen *net.TCPListener
	Host   string
	Port   int
}

// TCPServerFactory is global create a new TCPServer object
func TCPServerFactory(host string, port int) *TCPServer {
	return new(TCPServer).Factory(host, port)
}

// Factory is create a new TCPServer object
func (me *TCPServer) Factory(host string, port int) *TCPServer {
	me.Host = host
	me.Port = port
	return me
}

// Open is open connection of TCPServer
func (me *TCPServer) Open() (bool, error) {
	addr := fmt.Sprint(me.Host, ":", me.Port)
	rAddr, errRes := net.ResolveTCPAddr("tcp", addr)
	if !handler.Error(errRes) {
		l, errLis := net.ListenTCP("tcp", rAddr)
		if !handler.Error(errLis) {
			me.listen = l
			return true, errLis
		}
	}

	me.listen = nil
	return false, errRes
}

// Close is close connection of TCPServer
func (me *TCPServer) Close() (bool, error) {
	if me.listen != nil {
		errLis := me.listen.Close()
		if !handler.Error(errLis) {
			return true, errLis
		}
		return false, errLis
	}
	return false, errors.New("TCPServer Connection is not listen")
}

// Accept is accept a new connection and return TCPClient
func (me *TCPServer) Accept() (*TCPClient, error) {
	conn, errAcc := me.listen.AcceptTCP()
	if !handler.Error(errAcc) {
		client := new(TCPClient).Factory(me.host, me.port)
		client.conn = conn
		return client, errAcc
	}
	return nil, errAcc
}

// AcceptTimeout is accept with timeout in second
func (me *TCPServer) AcceptTimeout(sec int) (*TCPClient, error) {
	conn, errAcc := me.listen.AcceptTCP()
	if !handler.Error(errAcc) {
		ts := time.Second * time.Duration(sec)
		client := new(TCPClient).Factory(me.Host, me.Port)
		client.Conn = conn
		go TCPTimeOutRoutine(client, sec)
		client.conn.SetDeadline(time.Now().Add(ts))
		return client, errAcc
	}
	return nil, errAcc
}
