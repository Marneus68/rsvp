package spooler

import (
	"github.com/Marneus68/gvp/config"
	"log"
	"net"
	"time"
)

func Start(con *config.Config) {
	ln, err := net.Listen("tcp", con.Port)
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go spool(con, conn)
	}
}

func spool(con *config.Config, conn net.Conn) {
	err := conn.SetReadDeadline(time.Now().Add(time.Duration(con.Timeout) * time.Second))
	if err != nil {
		log.Fatal(err.Error())
	}
}
