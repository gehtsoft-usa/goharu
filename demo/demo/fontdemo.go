package demo

import (
	"github.com/gehtsoft-usa/goharu"
)

var fonts []string = []string{
	"Courier",
	"Courier-Bold",
	"Courier-Oblique",
	"Courier-BoldOblique",
	"Helvetica",
	"Helvetica-Bold",
	"Helvetica-Oblique",
	"Helvetica-BoldOblique",
	"Times-Roman",
	"Times-Bold",
	"Times-Italic",
	"Times-BoldItalic",
	"Symbol",
	"ZapfDingbats",
}

func Fontdemo() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()

	page := pdf.AddPage()
	height := page.GetHeight()
	width := page.GetWidth()

	page.SetLineWidth(1)
	page.Rectangle(50, 50, width-100, height-110)
	page.Stroke()

	defaultFont := pdf.GetFont("Helvetica", "")
	page.SetFontAndSize(defaultFont, 24)
	pageTitle := "Font Demo"
	tw := page.TextWidth(pageTitle)
	page.BeginText()
	page.TextOut((width-tw)/2, height-50, pageTitle)
	page.EndText()

	pageSubTitle := "Standard Type 1 font samples"
	page.SetFontAndSize(defaultFont, 16)
	page.BeginText()
	tw = page.TextWidth(pageSubTitle)
	page.TextOut((width-tw)/2, height-80, pageSubTitle)
	page.EndText()

	page.BeginText()
	page.MoveTextPos(60, height-105)

	sampleText := "abcdefgABCDEFG12345!#$%&+-@?"
	var i int
	for i = 0; i < len(fonts); i++ {
		font := pdf.GetFont(fonts[i], "")
		page.SetFontAndSize(defaultFont, 9)
		page.ShowText(fonts[i])
		page.MoveTextPos(0, -18)

		page.SetFontAndSize(font, 20)
		page.ShowText(sampleText)
		page.MoveTextPos(0, -20)
	}

	page.EndText()
	pdf.Save("out/fontdemo.pdf")
}
