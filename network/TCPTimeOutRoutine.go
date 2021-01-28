package network

import (
	"time"
)

// TCPTimeOutRoutine is thread for close connection
func TCPTimeOutRoutine(conn *TCPClient, sec int) (bool, error) {
	time.Sleep(time.Second * time.Duration(sec))
	return conn.Close()
}
