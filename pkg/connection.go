package pkg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

type Connection interface {
	Close() error
}

type connection struct {
	connection net.Conn
}

func Open(host string, port int) (Connection, error) {
	con := connection{}

	adr := fmt.Sprintf("%v:%v", host, port)

	if c, err := net.Dial("tcp", adr); err != nil {
		log.Error(err)
		return nil, err
	} else {
		con.connection = c
		return &con, nil
	}
}

func (r *connection) Close() error {
	return r.connection.Close()
}

/*
	message, err := bufio.NewReader(connection).ReadString('>')
			bytesWritten, err := fmt.Fprintf(connection, "slot-info\n")

*/
