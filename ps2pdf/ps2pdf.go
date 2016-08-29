package ps2pdf

import (
	"github.com/Marneus68/gvp/config"
	"log"
)

func Convert(
	inPath string,
	outDir string,
	con *config.Config,
) {
	log.Println("attempting to convert " + inPath + " to PDF")
}
