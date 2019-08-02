package demo

import (
	"github.com/gehtsoft-usa/goharu"
)

func Jpgdemo() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()

	pdf.SetCompressionMode(goharu.COMP_ALL)
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
	page.ShowText("jpgdemo")
	page.EndText()

	page.SetFontAndSize(font, 12)

	jpgDrawImage(pdf, "images/rgb.jpg", 70, page.GetHeight()-410, "24 bit color")
	jpgDrawImage(pdf, "images/gray.jpg", 340, page.GetHeight()-410, "grayscale")

	pdf.Save("out/jpgdemo.pdf")
}

func jpgDrawImage(pdf goharu.Doc, filename string, x float32, y float32, text string) {
	page := pdf.GetCurrentPage()

	image := pdf.LoadJpegImageFromFile(filename)

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
