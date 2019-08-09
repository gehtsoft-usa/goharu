package demo

import (
	"fmt"

	"github.com/gehtsoft-usa/goharu"
)

const ENC_PAGE_WIDTH float32 = 420
const ENC_PAGE_HEIGHT float32 = 400
const ENC_CELL_WIDTH int = 20
const ENC_CELL_HEIGHT int = 20
const ENC_CELL_HEADER int = 10

func Encodings() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()
	var encodings []string = []string{
		"StandardEncoding",
		"MacRomanEncoding",
		"WinAnsiEncoding",
		"ISO8859-2",
		"ISO8859-3",
		"ISO8859-4",
		"ISO8859-5",
		"ISO8859-9",
		"ISO8859-10",
		"ISO8859-13",
		"ISO8859-14",
		"ISO8859-15",
		"ISO8859-16",
		"CP1250",
		"CP1251",
		"CP1252",
		"CP1254",
		"CP1257",
		"KOI8-R",
		"Symbol-Set",
		"ZapfDingbats-Set",
	}

	pdf.SetCompressionMode(goharu.CompAll)
	pdf.SetPageMode(goharu.PageModeUseOutline)

	font := pdf.GetFont("Helvetica", "")

	fontName := pdf.LoadType1FontFromFile("type1/a010013l.afm", "type1/a010013l.pfb")

	/* create outline root. */
	root := pdf.CreateOutline(goharu.Outline{}, "Encoding list", goharu.Encoder{})
	root.SetOpened(true)

	var i int
	var page goharu.Page
	var outline goharu.Outline
	var font2 goharu.Font
	var dst goharu.Destination

	for i = 0; i < len(encodings); i++ {
		page = pdf.AddPage()
		page.SetWidth(ENC_PAGE_WIDTH)
		page.SetHeight(ENC_PAGE_HEIGHT)

		outline = pdf.CreateOutline(root, encodings[i], goharu.Encoder{})
		dst = page.CreateDestination()
		dst.SetXYZ(0, page.GetHeight(), 1)
		/* HPDF_Destination_SetFitB(dst); */
		outline.SetDestination(dst)

		page.SetFontAndSize(font, 15)
		drawGraph(page)

		page.BeginText()
		page.SetFontAndSize(font, 20)
		page.MoveTextPos(40, ENC_PAGE_HEIGHT-50)
		page.ShowText(encodings[i])
		page.ShowText(" Encoding")
		page.EndText()

		if encodings[i] == "Symbol-Set" {
			font2 = pdf.GetFont("Symbol", "")
		} else if encodings[i] == "ZapfDingbats-Set" {
			font2 = pdf.GetFont("ZapfDingbats", "")
		} else {
			font2 = pdf.GetFont(fontName, encodings[i])
		}

		page.SetFontAndSize(font2, 14)
		drawFonts(page)
	}

	pdf.Save("out/encodings.pdf")
}

func drawGraph(page goharu.Page) {
	var buf string
	var i, x, y int

	/* Draw 16 X 15 cells */

	/* Draw vertical lines. */
	page.SetLineWidth(0.5)

	for i = 0; i <= 17; i++ {
		x = i*ENC_CELL_WIDTH + 40

		page.MoveTo(float32(x), ENC_PAGE_HEIGHT-60)
		page.LineTo(float32(x), 40)
		page.Stroke()

		if i > 0 && i <= 16 {
			page.BeginText()
			page.MoveTextPos(float32(x+5), ENC_PAGE_HEIGHT-75)
			buf = fmt.Sprintf("%X", i-1)
			page.ShowText(buf)
			page.EndText()
		}
	}

	/* Draw horizontal lines. */
	for i = 0; i <= 15; i++ {
		y = i*ENC_CELL_HEIGHT + 40
		page.MoveTo(40, float32(y))
		page.LineTo(ENC_PAGE_WIDTH-40, float32(y))
		page.Stroke()

		if i < 14 {
			page.BeginText()
			page.MoveTextPos(45, float32(y+5))
			buf = fmt.Sprintf("%X", 15-i)
			page.ShowText(buf)
			page.EndText()
		}
	}
}

func drawFonts(page goharu.Page) {
	var i, j, x, y int
	var buff []byte = make([]byte, 2)
	var d float32

	buff[1] = 0
	page.BeginText()

	/* Draw all character from 0x20 to 0xFF to the canvas. */
	for i = 1; i < 17; i++ {
		for j = 1; j < 17; j++ {
			y = int(ENC_PAGE_HEIGHT) - 55 - ((i - 1) * ENC_CELL_HEIGHT)
			x = j*ENC_CELL_WIDTH + 50
			buff[0] = byte((i-1)*16 + (j - 1))
			if buff[0] >= 32 {
				d = float32(x) - page.TextWidth2(buff)/2
				page.TextOut2(float32(d), float32(y), buff)
			}
		}
	}
	page.EndText()
}
