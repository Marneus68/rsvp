package spooler

import (
	"github.com/Marneus68/gvp/config"
	"log"
	"net"
)

func Start(con *config.Config) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go spool(conn)
	}
}

func spool(conn net.Conn) {

}
