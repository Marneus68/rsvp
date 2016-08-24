package main

import (
	"flag"
	"github.com/Marneus68/gvp/config"
	"github.com/Marneus68/gvp/spooler"
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
	spooler.Start(config.NewConfig(outDir, port, mail, destMail, sendMail, smtp, smtpName, smtpPwd))
}
