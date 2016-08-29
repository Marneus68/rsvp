package spooler

import (
	"bufio"
	"github.com/Marneus68/gvp/config"
	"github.com/Marneus68/utils"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

func Start(con *config.Config) {
	// Create the output directory
	err := os.MkdirAll(utils.SubstituteHomeDir(con.OutDir), 0777)
	if err != nil {
		if os.IsExist(err) {
			log.Println(err.Error())
		} else {
			log.Panic(err.Error())
		}
	}
	ln, err := net.Listen("tcp", con.Port)
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go spool(con, conn, nil)
	}
}

func spool(
	con *config.Config,
	conn net.Conn,
	fun func(string, string, *config.Config),
) {
	defer func() {
		log.Println("Connection closed.")
		conn.Close()
	}()
	log.Print("New incomming connection...")
	err := conn.SetReadDeadline(time.Now().Add(time.Duration(con.Timeout) * time.Second))
	if err != nil {
		log.Println(err.Error())
	}
	ts := time.Now().Format(time.UnixDate)
	//bs, err := bufio.NewReader(conn).ReadBytes(io.EOF)
	scanner := bufio.NewScanner(conn)
	content := []byte("")
	for scanner.Scan() {
		log.Println("    " + scanner.Text())
		content = append(content, []byte(scanner.Text()+"\n")...)
	}

	out := utils.SubstituteHomeDir(con.OutDir) + string(filepath.Separator) + ts
	outPdl := out + ".pdl"

	ioutil.WriteFile(
		outPdl,
		[]byte(content),
		0777,
	)

	if err := scanner.Err(); err != nil {
		log.Print(err.Error())
	}

	if fun != nil {
		fun(outPdl, utils.SubstituteHomeDir(con.OutDir), con)
	}
}
