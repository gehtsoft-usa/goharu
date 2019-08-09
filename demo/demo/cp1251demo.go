package demo

import (
	"github.com/gehtsoft-usa/goharu"
)

func Cp1251Demo() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()

	//NOTE: libharu embed fonts have broken CP1251, the letters overlaps
	fontName := pdf.LoadTTFontFromFile("ttfont/a010013l.ttf", true)
	page := pdf.AddPage()
	font := pdf.GetFont(fontName, "CP1251")
	text, err := readFile("mbtext/cp1251.txt")
	if err != nil {
		text = []byte("Can't read text")
	}
	page.SetSize(goharu.PageSizeLetter, goharu.PageLandscape)
	page.BeginText()
	page.SetFontAndSize(font, 10)
	tw := page.TextWidth2(text)
	page.MoveTextPos((page.GetWidth()-tw)/2, (page.GetHeight()-20)/2)
	page.ShowText2(text)
	page.EndText()

	page = pdf.AddPage()
	text1 := "Травка зеленеет, Солнышко блестит, Ласточка с весною, В гости к нам летит"
	page.SetSize(goharu.PageSizeLegal, goharu.PageLandscape)
	page.BeginText()
	page.SetFontAndSize(font, 10)
	tw, _ = page.TextWidth1(text1, "CP1251")
	page.MoveTextPos((page.GetWidth()-tw)/2, (page.GetHeight()-20)/2)
	page.ShowText1(text1, "CP1251")
	page.EndText()

	pdf.Save("out/cp1251.pdf")
}
