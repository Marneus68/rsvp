package ps2pdf

import (
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dpdf"
	"github.com/llgcode/ps"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Convert a postscript file to a pdf file
func Convert(
	inPath string,
	outDir string,
	endCb func(string),
) {
	log.Println("Attempting to convert " + inPath + " to PDF")

	// get the file name without the extention
	fn := filepath.Base(inPath)
	sfn := fn[:len(fn)-len(filepath.Ext(fn))]
	// create the output path
	outPath := filepath.Join(outDir, sfn) + ".pdf"

	// prepare the pdf's graphic context
	pdf := draw2dpdf.NewPdf("P", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(pdf)

	DrawPsInGc(gc, inPath)
	//gc.Restore()

	log.Println("New PDF file will be created at " + outPath)

	draw2dpdf.SaveToPdfFile(outPath, pdf)

	if endCb != nil {
		endCb(inPath[:len(inPath)-len(filepath.Ext(inPath))])
	}
}

// Draw the content of a postscript file in a GraphicContext
func DrawPsInGc(gc draw2d.GraphicContext, filename string) {
	// Open the postscript
	src, err := os.OpenFile(filename, 0, 0)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer src.Close()
	bytes, err := ioutil.ReadAll(src)
	reader := strings.NewReader(string(bytes))

	// Initialize and interpret the postscript
	interpreter := ps.NewInterpreter(gc)
	interpreter.Execute(reader)
}
