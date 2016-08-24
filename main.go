package main

import (
	"flag"
	//"fmt"
	"github.com/Marneus68/gvp/config"
	"github.com/Marneus68/gvp/spooler"
	"github.com/Marneus68/utils"
	"log"
)

var outDir string
var port string

var mail bool

var destMail string
var sendMail string
var smtp string
var smtpName string
var smtpPwd string

func init() {
	flag.StringVar(&outDir, "o", "~/gvp_tray", "output directory")
	flag.StringVar(&port, "p", ":9100", "port gvp listens on")
	flag.BoolVar(&mail, "m", false, "send prints via email")
	flag.StringVar(&destMail, "dm", "user@cock.li", "destination email address")
	flag.StringVar(&sendMail, "sm", "gvp-noreply@hostname", "sender email address")
	flag.StringVar(&smtp, "sa", "mail.cock.li:587", "smtp server and port")
	flag.StringVar(&smtpName, "su", "username@domail.tld", "smtp user name")
	flag.StringVar(&smtpPwd, "sp", "password", "smtp user password")
}

func main() {
	flag.Parse()
	/*
		fmt.Println(outputDir)
		fmt.Println(port)
	*/
	valid, normalPort := utils.IsValidPortString(port)
	if !valid {
		log.Fatal("Invalid port string")
	}
	con := config.NewConfig(outDir, normalPort, mail, destMail, sendMail, smtp, smtpName, smtpPwd)
	spooler.Start(con)
}
