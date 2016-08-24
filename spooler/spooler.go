package spooler

import (
	"github.com/Marneus68/gvp/config"
	"net"
)

func Start(con *config.Config) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go spool(conn)
	}
}

func spool(conn net.Conn) {

}
