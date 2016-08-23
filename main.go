package main

import (
	"flag"
	"fmt"
)

var outputDir string
var port string

var sendMail bool

var emailAddress string
var smtp string
var smtpName string
var smtpPwd string

func init() {
	flag.StringVar(&outputDir, "o", "~/gvp_tray", "output directory")
	flag.StringVar(&port, "p", ":9100", "port gvp listens on")
	flag.BoolVar(&sendMail, "m", false, "send prints via email")
	flag.StringVar(&port, "a", ":9100", "email address")
	flag.StringVar(&smtp, "s", ":9100", "smtp server and port")
	flag.StringVar(&smtpName, "su", ":9100", "smtp user name")
	flag.StringVar(&smtpPwd, "sp", ":9100", "smtp user password")
}

func main() {
	flag.Parse()
	fmt.Println(outputDir)
	fmt.Println(port)
}
