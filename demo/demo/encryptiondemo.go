package demo

import (
	"github.com/gehtsoft-usa/goharu"
)

func Encryptiondemo() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()

	page := pdf.AddPage()
	font := pdf.GetFont("Helvetica", "")
	text := "This is an encrypted document example."
	page.SetSize(goharu.PageSizeB5, goharu.PageLandscape)
	page.BeginText()
	page.SetFontAndSize(font, 20)
	tw := page.TextWidth(text)
	page.MoveTextPos((page.GetWidth()-tw)/2, (page.GetHeight()-20)/2)
	page.ShowText(text)
	page.EndText()
	pdf.SetPassword("owner", "user")
	pdf.Save("out/encryptiondemo.pdf")
}
