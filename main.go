package main

import (
	"flag"
	"fmt"
)

var outputDir string
var port string

func init() {
	flag.StringVar(&outputDir, "o", "~/gvp", "output directory")
	flag.StringVar(&port, "p", ":9100", "port gvp listens on")
}

func main() {
	flag.Parse()
	fmt.Println(outputDir)
	fmt.Println(port)
}
