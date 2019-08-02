package demo

import (
	"github.com/gehtsoft-usa/goharu"
)

func JpDemo() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()

	page := pdf.AddPage()
	pdf.UseJPFonts()
	pdf.UseJPEncodings()

	fontName := pdf.LoadTTFontFromFile("ttfont/chi_jyun.ttf", true)

	font := pdf.GetFont(fontName, "90ms-RKSJ-H")
	text, err := readFile("mbtext/cp932.txt")
	if err != nil {
		text = []byte("Can't read text")
	}
	page.SetSize(goharu.PAGE_SIZE_LETTER, goharu.PAGE_LANDSCAPE)
	page.BeginText()
	page.SetFontAndSize(font, 10)
	tw := page.TextWidth2(text)
	page.MoveTextPos((page.GetWidth()-tw)/2, (page.GetHeight()-20)/2)
	page.ShowText2(text)
	page.EndText()

	pdf.Save("out/jp.pdf")
}
