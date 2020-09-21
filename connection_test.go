package connection_test

import (
	"testing"

	"github.com/ifrasoft/connection"
	"github.com/magiconair/properties/assert"
)

func TestAddConnection(t *testing.T) {
	cl := connection.NewConnectionLog()
	cl.AddConnection(&connection.Connection{Name: "FTP", Host: "127.0.0.1", Port: "21"})
	cl.AddConnection(&connection.Connection{Name: "HTTP", Host: "127.0.0.1", Port: "80"})
	cl.AddConnection(&connection.Connection{Name: "HTTPS", Host: "127.0.0.1", Port: "443"})
	assert.Equal(t, len(cl.GetConnections()), 3)
}
