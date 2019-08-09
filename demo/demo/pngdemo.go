package demo

import (
	"github.com/gehtsoft-usa/goharu"
)

func Pngdemo() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()

	pdf.SetCompressionMode(goharu.CompAll)
	font := pdf.GetFont("Helvetica", "")

	page := pdf.AddPage()
	page.SetWidth(650)
	page.SetHeight(500)

	dst := page.CreateDestination()
	dst.SetXYZ(225, page.GetHeight(), 1)
	pdf.SetOpenAction(dst)

	page.BeginText()
	page.SetFontAndSize(font, 20)
	page.MoveTextPos(220, page.GetHeight()-70)
	page.ShowText("pngdemo")
	page.EndText()

	page.SetFontAndSize(font, 8)

	//uncomment 16bit color as soon as fixed static libraries are available
	pngDrawImage(pdf, "pngsuite/basn0g01.png", 100, page.GetHeight()-150, "1bit grayscale.")
	pngDrawImage(pdf, "pngsuite/basn0g02.png", 200, page.GetHeight()-150, "2bit grayscale.")
	pngDrawImage(pdf, "pngsuite/basn0g04.png", 300, page.GetHeight()-150, "4bit grayscale.")
	pngDrawImage(pdf, "pngsuite/basn0g08.png", 400, page.GetHeight()-150, "8bit grayscale.")
	pngDrawImage(pdf, "pngsuite/basn2c08.png", 100, page.GetHeight()-250, "8bit color.")
	//pngDrawImage(pdf, "pngsuite/basn2c16.png", 200, page.GetHeight()-250, "16bit color.")
	pngDrawImage(pdf, "pngsuite/basn3p01.png", 100, page.GetHeight()-350, "1bit palette.")
	pngDrawImage(pdf, "pngsuite/basn3p02.png", 200, page.GetHeight()-350, "2bit palette.")
	pngDrawImage(pdf, "pngsuite/basn3p04.png", 300, page.GetHeight()-350, "4bit palette.")
	pngDrawImage(pdf, "pngsuite/basn3p08.png", 400, page.GetHeight()-350, "8bit palette.")
	pngDrawImage(pdf, "pngsuite/basn4a08.png", 100, page.GetHeight()-450, "8bit alpha.")
	//pngDrawImage(pdf, "pngsuite/basn4a16.png", 200, page.GetHeight()-450, "16bit alpha.")
	pngDrawImage(pdf, "pngsuite/basn6a08.png", 100, page.GetHeight()-550, "8bit alpha.")
	//pngDrawImage(pdf, "pngsuite/basn6a16.png", 200, page.GetHeight()-550, "16bit alpha.")

	pdf.Save("out/pngdemo.pdf")
}

func pngDrawImage(pdf goharu.Doc, filename string, x float32, y float32, text string) {
	page := pdf.GetCurrentPage()

	image := pdf.LoadPngImageFromFile(filename)

	/* Draw image to the canvas. */
	page.DrawImage(image, x, y, float32(image.Width()), float32(image.Height()))

	/* Print the text. */
	page.BeginText()
	page.SetTextLeading(16)
	page.MoveTextPos(x, y)
	page.ShowTextNextLine(filename)
	page.ShowTextNextLine(text)
	page.EndText()
}
