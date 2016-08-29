package main

import (
	"flag"
	"github.com/Marneus68/rsvp/config"
	"github.com/Marneus68/rsvp/spooler"
)

var outDir string
var port string

var mail bool

var destMail string
var sendMail string
var smtp string
var smtpName string
var smtpPwd string

var timeout int

func init() {
	flag.StringVar(&outDir, "o", "~/rsvp", "output directory")
	flag.StringVar(&port, "p", ":9100", "port rsvp listens on")
	flag.BoolVar(&mail, "m", false, "send prints via email")
	flag.StringVar(&destMail, "dm", "user@domain.tld", "destination email address")
	flag.StringVar(&sendMail, "sm", "rsvp-noreply@hostname", "sender email address")
	flag.StringVar(&smtp, "sa", "smtp.domain.tld:587", "smtp server and port")
	flag.StringVar(&smtpName, "su", "user@domain.tld", "smtp user name")
	flag.StringVar(&smtpPwd, "sp", "password", "smtp user password")
	flag.IntVar(&timeout, "t", 5, "timeout in seconds")
}

func main() {
	flag.Parse()
	spooler.Start(config.NewConfig(outDir, port, mail, destMail, sendMail, smtp, smtpName, smtpPwd, timeout))
}
