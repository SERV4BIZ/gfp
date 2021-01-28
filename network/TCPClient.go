package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net"
	"time"

	"github.com/SERV4BIZ/gfp/compress"
	"github.com/SERV4BIZ/gfp/handler"
)

// TCPClient is structure
type TCPClient struct {
	Conn *net.TCPConn
	Host string
	Port int
}

// TCPClientFactory is global create a new TCPClient object
func TCPClientFactory(host string, port int) *TCPClient {
	return new(TCPClient).Factory(host, port)
}

// Factory is create a new TCPClient object
func (me *TCPClient) Factory(host string, port int) *TCPClient {
	me.Conn = nil
	me.Host = host
	me.Port = port
	return me
}

// Connect is open new connection
func (me *TCPClient) Connect() (bool, error) {
	addr := fmt.Sprint(me.Host, ":", me.Port)
	rAddr, errAddr := net.ResolveTCPAddr("tcp", addr)
	if !handler.Error(errAddr) {
		conn, errDial := net.DialTCP("tcp", nil, rAddr)
		if !handler.Error(errDial) {
			me.Conn = conn
			return true, errDial
		}
		return false, errDial
	}

	me.Conn = nil
	return false, errAddr
}

// ConnectTimeout is open new connection with timeout
func (me *TCPClient) ConnectTimeout(sec int) (bool, error) {
	addr := fmt.Sprint(me.Host, ":", me.Port)
	ts := time.Second * time.Duration(sec)
	conn, errDial := net.DialTimeout("tcp", addr, ts)
	if !handler.Error(errDial) {
		me.Conn = conn.(*net.TCPConn)
		go TCPTimeOutRoutine(me, sec)
		me.Conn.SetDeadline(time.Now().Add(ts))
		return true, errDial
	}
	me.Conn = nil
	return false, errDial
}

// Close is close connection
func (me *TCPClient) Close() (bool, error) {
	if me.Conn != nil {
		errClose := me.Conn.Close()
		return errClose == nil, errClose
	}
	return false, errors.New("TCPClient Connection is not open")
}

// IntToByte is convert int to byte buffer
func (me *TCPClient) IntToByte(value int) []byte {
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, uint32(value))
	return buff
}

// ByteToInt is convert byte buffer to int
func (me *TCPClient) ByteToInt(buffer []byte) (int, error) {
	var size int32
	errRead := binary.Read(bytes.NewReader(buffer), binary.BigEndian, &size)
	if !handler.Error(errRead) {
		return int(size), errRead
	}
	return -1, errRead
}

// Send is send data byte buffer
func (me *TCPClient) Send(buffer []byte) (int, error) {
	if me.Conn != nil {
		size, errWrite := me.Conn.Write(buffer)
		if !handler.Error(errWrite) {
			return size, errWrite
		}
		return -1, errWrite
	}
	return -1, errors.New("TCPClient Connection is not open")
}

//Receive is get data byte buffer
func (me *TCPClient) Receive(buffer []byte) (int, error) {
	if me.Conn != nil {
		size, errRead := me.Conn.Read(buffer)
		if !handler.Error(errRead) {
			return size, errRead
		}
		return -1, errRead
	}
	return -1, errors.New("TCPClient Connection is not open")
}

// SendByte is send data byte buffer
func (me *TCPClient) SendByte(buffer []byte) (int, error) {
	size := binary.Size(buffer)
	_, errSend := me.Send(me.IntToByte(size))
	if !handler.Error(errSend) {
		_, err := me.Send(buffer)
		if !handler.Error(err) {
			return size, err
		}
	}
	return -1, errSend
}

// ReceiveByte is get byte data buffer
func (me *TCPClient) ReceiveByte() ([]byte, error) {
	bsizes := make([]byte, 4)
	_, errRec := me.Receive(bsizes)
	if !handler.Error(errRec) {
		size, errBTI := me.ByteToInt(bsizes)
		if !handler.Error(errBTI) {
			if size < 0 || size > math.MaxInt64 {
				size = 0
			}
			buffers := make([]byte, size)
			getbyte := make([]byte, 1)
			var resulterr error
			for i := 0; i < size; i++ {
				_, err := me.Receive(getbyte)
				if !handler.Error(err) {
					buffers[i] = getbyte[0]
				} else {
					resulterr = err
					break
				}
			}
			return buffers, resulterr
		}
		return nil, errBTI
	}
	return nil, errRec
}

// SendString is send string buffer
func (me *TCPClient) SendString(buffer string) (int, error) {
	return me.SendByte([]byte(buffer))
}

// ReceiveString is get string buffer
func (me *TCPClient) ReceiveString() (string, error) {
	buffer, errRB := me.ReceiveByte()
	if !handler.Error(errRB) {
		return string(buffer), errRB
	}
	return "", errRB
}

// SendDeflate is send byte buffer with compress
func (me *TCPClient) SendDeflate(buffer []byte) (int, error) {
	b, errComp := compress.Encode(buffer)
	if !handler.Error(errComp) {
		return me.SendByte(b)
	}
	return -1, errComp
}

// ReceiveDeflate is get byte buffer with decompress
func (me *TCPClient) ReceiveDeflate() ([]byte, error) {
	b, errRec := me.ReceiveByte()
	if !handler.Error(errRec) {
		return compress.Decode(b)
	}
	return nil, errRec
}
