package ps2pdf

import (
	"log"
	"path/filepath"
)

func Convert(
	inPath string,
	outDir string,
	endCb func(string),
) {
	log.Println("attempting to convert " + inPath + " to PDF")

	if endCb != nil {
		endCb(inPath[:len(inPath)-len(filepath.Ext(inPath))])
	}
}
