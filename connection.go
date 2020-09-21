package connection

import (
	"net"
	"time"
)

type Connection struct {
	Name   string
	Host   string
	Port   string
	Status string
}

type ConnectionLog interface {
	AddConnection(con *Connection)
	Check() []*Result
	GetConnections() []*Connection
}

type Result struct {
	Connection *Connection
	Status     string
	Error      error
}

type connectionLog struct {
	Connections []*Connection
}

func NewConnectionLog() ConnectionLog {
	return &connectionLog{}
}

func (cl *connectionLog) AddConnection(con *Connection) {
	cl.Connections = append(cl.Connections, con)
}

func (cl *connectionLog) Check() []*Result {
	results := []*Result{}
	ch := make(chan Result, 4)

	for _, connection := range cl.Connections {
		go cl.CheckConnect(connection, ch)
	}

	count := 0
	for {
		result := <-ch
		results = append(results, &result)
		count++
		if count == len(cl.Connections) {
			break
		}
	}

	return results
}

func (cl *connectionLog) GetConnections() []*Connection {
	return cl.Connections
}

func (cl *connectionLog) CheckConnect(con *Connection, rs chan Result) {

	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(con.Host, con.Port), timeout)

	if err != nil {
		rs <- Result{
			Error:      err,
			Status:     "FAIL",
			Connection: con,
		}

	}

	if conn != nil {
		defer conn.Close()

	}

	rs <- Result{
		Error:      err,
		Status:     "OK",
		Connection: con,
	}

}
