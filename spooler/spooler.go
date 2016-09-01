package spooler

import (
	"bufio"
	"github.com/Marneus68/rsvp/config"
	"github.com/Marneus68/rsvp/ps2pdf"
	"github.com/Marneus68/utils"
	"github.com/go-gomail/gomail"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
		go spool(con, conn, ps2pdf.Convert, func(dest string) {
			if con.Mail == true {

				smtps := strings.Split(con.Smtp, ":")
				if len(smtps) != 2 {
					log.Println("Invalid smtp string format")
				}

				smtpp, err := strconv.Atoi(smtps[1])
				if err != nil {
					log.Println(err.Error())
				}

				m := gomail.NewMessage()
				m.SetHeader("From", con.SendMail)
				m.SetHeader("To", con.DestMail)
				m.SetHeader("Subject", "New print received on your virtual printer")
				m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
				m.Attach("/home/Alex/lolcat.jpg")

				d := gomail.NewDialer(smtps[0], smtpp, con.SmtpName, con.SmtpPwd)

				// Send the email
				if err := d.DialAndSend(m); err != nil {
					log.Println(err.Error())
				}
			}
		})
	}
}

func spool(
	con *config.Config,
	conn net.Conn,
	psFun func(string, string, func(string)),
	mailFun func(string),
) {
	defer func() {
		log.Println("Connection closed.")
		conn.Close()
	}()
	log.Println("New incomming connection...")
	err := conn.SetReadDeadline(time.Now().Add(time.Duration(con.Timeout) * time.Second))
	if err != nil {
		log.Println(err.Error())
	}
	ts := time.Now().Format(time.UnixDate)
	scanner := bufio.NewScanner(conn)
	content := []byte("")
	for scanner.Scan() {
		content = append(content, []byte(scanner.Text()+"\n")...)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	out := utils.SubstituteHomeDir(con.OutDir) + string(filepath.Separator) + ts
	pdl := filepath.FromSlash(filepath.Clean(out + ".pdl"))

	err = ioutil.WriteFile(
		pdl,
		[]byte(content),
		0666,
	)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Created temporary file in " + pdl)
	if psFun != nil {
		psFun(pdl, utils.SubstituteHomeDir(con.OutDir), mailFun)
	}
}
