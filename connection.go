package connection

type Connection struct {
	Name   string
	Host   string
	Port   string
	Status string
}

type ConnectionLog interface {
	AddConnection(con *Connection)
	Check()
	GetConnections() []*Connection
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

func (cl *connectionLog) Check() {

}

func (cl *connectionLog) GetConnections() []*Connection {
	return cl.Connections
}
