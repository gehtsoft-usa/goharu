package demo

import (
	"fmt"
	"math"

	"github.com/gehtsoft-usa/goharu"
)

func Textdemo() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()

	page := pdf.AddPage()

	sampleText := "abcdefgABCDEFG123!#$%&+-@?"
	sampleText2 := "The quick brown fox jumps over the lazy dog."
	pageTitle := "Font Sample"
	font := pdf.GetFont("Helvetica", "")

	var i, length uint
	var tw, fsize, ypos float32
	var angle1, angle2, rad1, rad2 float64
	var buf string

	printGrid(pdf, page)

	/* print the title of the page (with positioning center). */
	page.SetFontAndSize(font, 24)
	tw = page.TextWidth(pageTitle)
	page.BeginText()
	page.TextOut((page.GetWidth()-tw)/2, page.GetHeight()-50, pageTitle)
	page.EndText()

	page.BeginText()
	page.MoveTextPos(60, page.GetHeight()-60)

	/*
	 * font size
	 */
	fsize = 8
	for fsize < 60 {
		/* set style and size of font. */
		page.SetFontAndSize(font, fsize)

		/* set the position of the text. */
		page.MoveTextPos(0, -5-fsize)

		/* measure the number of characters which included in the page. */
		length, _ = page.MeasureText(sampleText, page.GetWidth()-120, false)

		/* truncate the text. */
		buf = sampleText[:length]
		page.ShowText(buf)

		/* print the description. */
		page.MoveTextPos(0, -10)
		page.SetFontAndSize(font, 8)
		buf = fmt.Sprintf("Fontsize=%.0f", fsize)
		page.ShowText(buf)
		fsize *= 1.5
	}

	/*
	 * font color
	 */
	page.SetFontAndSize(font, 8)
	page.MoveTextPos(0, -30)
	page.ShowText("Font color")

	page.SetFontAndSize(font, 18)
	page.MoveTextPos(0, -20)
	length = uint(len(sampleText))
	for i = 0; i < length; i++ {

		r := float32(i) / float32(length)
		g := 1 - float32(i)/float32(length)
		buf = sampleText[i : i+1]
		page.SetRGBFill(goharu.RGBColor{R: r, G: g, B: 0.0})
		page.ShowText(buf)
	}
	page.MoveTextPos(0, -25)

	for i = 0; i < length; i++ {
		r := float32(i) / float32(length)
		b := 1 - float32(i)/float32(length)
		buf = sampleText[i : i+1]

		page.SetRGBFill(goharu.RGBColor{R: r, G: 0.0, B: b})
		page.ShowText(buf)
	}
	page.MoveTextPos(0, -25)

	for i = 0; i < length; i++ {
		b := float32(i) / float32(length)
		g := 1 - float32(i)/float32(length)
		buf = sampleText[i : i+1]

		page.SetRGBFill(goharu.RGBColor{R: 0.0, G: g, B: b})
		page.ShowText(buf)
	}

	page.EndText()

	ypos = 450

	/*
	 * Font rendering mode
	 */
	page.SetFontAndSize(font, 32)
	page.SetRGBFill(goharu.RGBColor{R: 0.5, G: 0.5, B: 0.0})
	page.SetLineWidth(1.5)

	/* PDF_FILL */
	showDescription(page, 60, ypos, "RenderingMode=PDF_FILL")
	page.SetTextRenderingMode(goharu.TEXT_RENDER_FILL)
	page.BeginText()
	page.TextOut(60, ypos, "ABCabc123")
	page.EndText()

	/* PDF_STROKE */
	showDescription(page, 60, ypos-50, "RenderingMode=PDF_STROKE")
	page.SetTextRenderingMode(goharu.TEXT_RENDER_STROKE)
	page.BeginText()
	page.TextOut(60, ypos-50, "ABCabc123")
	page.EndText()

	/* PDF_FILL_THEN_STROKE */
	showDescription(page, 60, ypos-100, "RenderingMode=PDF_FILL_THEN_STROKE")
	page.SetTextRenderingMode(goharu.TEXT_RENDER_FILL_THEN_STROKE)
	page.BeginText()
	page.TextOut(60, ypos-100, "ABCabc123")
	page.EndText()

	/* PDF_FILL_CLIPPING */
	showDescription(page, 60, ypos-150, "RenderingMode=PDF_FILL_CLIPPING")
	page.GSave()
	page.SetTextRenderingMode(goharu.TEXT_RENDER_FILL_CLIPPING)
	page.BeginText()
	page.TextOut(60, ypos-150, "ABCabc123")
	page.EndText()
	showStripePattern(page, 60, ypos-150)
	page.GRestore()

	/* PDF_STROKE_CLIPPING */
	showDescription(page, 60, ypos-200, "RenderingMode=PDF_STROKE_CLIPPING")
	page.GSave()
	page.SetTextRenderingMode(goharu.TEXT_RENDER_STROKE_CLIPPING)
	page.BeginText()
	page.TextOut(60, ypos-200, "ABCabc123")
	page.EndText()
	showStripePattern(page, 60, ypos-200)
	page.GRestore()

	/* PDF_FILL_STROKE_CLIPPING */
	showDescription(page, 60, ypos-250, "RenderingMode=PDF_FILL_STROKE_CLIPPING")
	page.GSave()
	page.SetTextRenderingMode(goharu.TEXT_RENDER_FILL_STROKE_CLIPPING)
	page.BeginText()
	page.TextOut(60, ypos-250, "ABCabc123")
	page.EndText()
	showStripePattern(page, 60, ypos-250)
	page.GRestore()

	/* Reset text attributes */
	page.SetTextRenderingMode(goharu.TEXT_RENDER_FILL)
	page.SetRGBFill(goharu.RGBColor{R: 0, G: 0, B: 0})
	page.SetFontAndSize(font, 30)

	/*
	 * Rotating text
	 */
	angle1 = 30                    /* A rotation of 30 degrees. */
	rad1 = angle1 / 180 * 3.141592 /* Calcurate the radian value. */

	showDescription(page, 320, ypos-60, "Rotating text")
	page.BeginText()
	page.SetTextMatrix(goharu.TransMatrix{A: float32(math.Cos(rad1)), B: float32(math.Sin(rad1)), C: float32(-math.Sin(rad1)), D: float32(math.Cos(rad1)), X: 330, Y: ypos - 60})
	page.ShowText("ABCabc123")
	page.EndText()

	/*
	 * Skewing text.
	 */
	showDescription(page, 320, ypos-120, "Skewing text")
	page.BeginText()

	angle1 = 10
	angle2 = 20
	rad1 = angle1 / 180 * 3.141592
	rad2 = angle2 / 180 * 3.141592

	page.SetTextMatrix(goharu.TransMatrix{A: 1, B: float32(math.Tan(rad1)), C: float32(math.Tan(rad2)), D: 1, X: 320, Y: ypos - 120})
	page.ShowText("ABCabc123")
	page.EndText()

	/*
	 * scaling text (X direction)
	 */
	showDescription(page, 320, ypos-175, "Scaling text (X direction)")
	page.BeginText()
	page.SetTextMatrix(goharu.TransMatrix{A: 1.5, B: 0, C: 0, D: 1, X: 320, Y: ypos - 175})
	page.ShowText("ABCabc12")
	page.EndText()

	/*
	 * scaling text (Y direction)
	 */
	showDescription(page, 320, ypos-250, "Scaling text (Y direction)")
	page.BeginText()
	page.SetTextMatrix(goharu.TransMatrix{A: 1, B: 0, C: 0, D: 2, X: 320, Y: ypos - 250})
	page.ShowText("ABCabc123")
	page.EndText()

	/*
	 * char spacing, word spacing
	 */

	showDescription(page, 60, 140, "char-spacing 0")
	showDescription(page, 60, 100, "char-spacing 1.5")
	showDescription(page, 60, 60, "char-spacing 1.5, word-spacing 2.5")

	page.SetFontAndSize(font, 20)
	page.SetRGBFill(goharu.RGBColor{R: 0.1, G: 0.3, B: 0.1})

	/* char-spacing 0 */
	page.BeginText()
	page.TextOut(60, 140, sampleText2)
	page.EndText()

	/* char-spacing 1.5 */
	page.SetCharSpace(1.5)

	page.BeginText()
	page.TextOut(60, 100, sampleText2)
	page.EndText()

	/* char-spacing 1.5, word-spacing 3.5 */
	page.SetWordSpace(2.5)

	page.BeginText()
	page.TextOut(60, 60, sampleText2)
	page.EndText()

	page = pdf.AddPage()
	page.SetSize(goharu.PAGE_SIZE_A5, goharu.PAGE_PORTRAIT)
	printGrid(pdf, page)

	/* text_rect method */

	/* ALIGN_LEFT */
	rect := goharu.Rect{}
	rect.Left = 25
	rect.Top = 545
	rect.Right = 200
	rect.Bottom = rect.Top - 40

	page.Rectangle(rect.Left, rect.Bottom, rect.Right-rect.Left, rect.Top-rect.Bottom)
	page.Stroke()
	page.BeginText()
	page.SetFontAndSize(font, 10)
	page.TextOut(rect.Left, rect.Top+3, "TEXT_ALIGN_LEFT")
	page.SetFontAndSize(font, 13)
	page.TextRect(rect.Left, rect.Top, rect.Right, rect.Bottom, sampleText2, goharu.TEXT_ALIGN_LEFT)
	page.EndText()

	/* ALIGN_RIGHT */
	rect.Left = 220
	rect.Right = 395
	page.Rectangle(rect.Left, rect.Bottom, rect.Right-rect.Left, rect.Top-rect.Bottom)
	page.Stroke()
	page.BeginText()
	page.SetFontAndSize(font, 10)
	page.TextOut(rect.Left, rect.Top+3, "TEXT_ALIGN_RIGHT")

	page.SetFontAndSize(font, 13)
	page.TextRect(rect.Left, rect.Top, rect.Right, rect.Bottom, sampleText2, goharu.TEXT_ALIGN_RIGHT)
	page.EndText()

	/* ALIGN_CENTER */
	rect.Left = 25
	rect.Top = 475
	rect.Right = 200
	rect.Bottom = rect.Top - 40

	page.Rectangle(rect.Left, rect.Bottom, rect.Right-rect.Left, rect.Top-rect.Bottom)
	page.Stroke()
	page.BeginText()
	page.SetFontAndSize(font, 10)
	page.TextOut(rect.Left, rect.Top+3, "TEXT_ALIGN_CENTER")

	page.SetFontAndSize(font, 13)
	page.TextRect(rect.Left, rect.Top, rect.Right, rect.Bottom, sampleText2, goharu.TEXT_ALIGN_CENTER)
	page.EndText()

	/* ALIGN_JUSTIFY */
	rect.Left = 220
	rect.Right = 395

	page.Rectangle(rect.Left, rect.Bottom, rect.Right-rect.Left, rect.Top-rect.Bottom)
	page.Stroke()
	page.BeginText()
	page.SetFontAndSize(font, 10)
	page.TextOut(rect.Left, rect.Top+3, "TEXT_ALIGN_JUSTIFY")

	page.SetFontAndSize(font, 13)
	page.TextRect(rect.Left, rect.Top, rect.Right, rect.Bottom, sampleText2, goharu.TEXT_ALIGN_JUSTIFY)
	page.EndText()

	/* Skewed coordinate system */
	page.GSave()

	angle1 = 5
	angle2 = 10
	rad1 = angle1 / 180 * 3.141592
	rad2 = angle2 / 180 * 3.141592

	page.Concat(1, float32(math.Tan(rad1)), float32(math.Tan(rad2)), 1, 25, 350)
	rect.Left = 0
	rect.Top = 40
	rect.Right = 175
	rect.Bottom = 0

	page.Rectangle(rect.Left, rect.Bottom, rect.Right-rect.Left, rect.Top-rect.Bottom)
	page.Stroke()
	page.BeginText()
	page.SetFontAndSize(font, 10)
	page.TextOut(rect.Left, rect.Top+3, "Skewed coordinate system")
	page.SetFontAndSize(font, 13)
	page.TextRect(rect.Left, rect.Top, rect.Right, rect.Bottom, sampleText2, goharu.TEXT_ALIGN_LEFT)
	page.EndText()
	page.GRestore()

	/* Rotated coordinate system */
	page.GSave()
	angle1 = 5
	rad1 = angle1 / 180 * 3.141592
	page.Concat(float32(math.Cos(rad1)), float32(math.Sin(rad1)), float32(-math.Sin(rad1)), float32(math.Cos(rad1)), 220, 350)

	rect.Left = 0
	rect.Top = 40
	rect.Right = 175
	rect.Bottom = 0

	page.Rectangle(rect.Left, rect.Bottom, rect.Right-rect.Left, rect.Top-rect.Bottom)
	page.Stroke()
	page.BeginText()
	page.SetFontAndSize(font, 10)
	page.TextOut(rect.Left, rect.Top+3, "Rotated coordinate system")
	page.SetFontAndSize(font, 13)
	page.TextRect(rect.Left, rect.Top, rect.Right, rect.Bottom, sampleText2, goharu.TEXT_ALIGN_LEFT)
	page.EndText()
	page.GRestore()

	/* text along a circle */
	page.SetGrayStroke(0)
	page.Circle(210, 190, 145)
	page.Circle(210, 190, 113)
	page.Stroke()

	angle1 = 360.0 / float64(len(sampleText2))
	angle2 = 180

	page.BeginText()
	page.SetFontAndSize(font, 30)

	for i = 0; i < uint(len(sampleText2)); i++ {
		rad1 = (angle2 - 90) / 180 * 3.141592
		rad2 = angle2 / 180 * 3.141592

		x := 210 + math.Cos(rad2)*122
		y := 190 + math.Sin(rad2)*122

		page.SetTextMatrix(goharu.TransMatrix{A: float32(math.Cos(rad1)), B: float32(math.Sin(rad1)), C: float32(-math.Sin(rad1)), D: float32(math.Cos(rad1)), X: float32(x), Y: float32(y)})

		buf = sampleText2[i : i+1]
		page.ShowText(buf)
		angle2 -= angle1
	}

	page.SetTextLeading(20)

	pdf.Save("out/textdemo.pdf")
}

func showStripePattern(page goharu.Page, x float32, y float32) {
	var iy int = 0

	for iy < 50 {
		page.SetRGBStroke(goharu.RGBColor{R: 0.0, G: 0.0, B: 0.5})
		page.SetLineWidth(1)
		page.MoveTo(x, y+float32(iy))
		page.LineTo(x+page.TextWidth("ABCabc123"), y+float32(iy))
		page.Stroke()
		iy += 3
	}

	page.SetLineWidth(2.5)
}

func showDescription(page goharu.Page, x float32, y float32, text string) {
	fsize := page.GetCurrentFontSize()
	font := page.GetCurrentFont()
	c := page.GetRGBFill()

	page.BeginText()
	page.SetRGBFill(goharu.RGBColor{R: 0, G: 0, B: 0})
	page.SetTextRenderingMode(goharu.TEXT_RENDER_FILL)
	page.SetFontAndSize(font, 10)
	page.TextOut(x, y-12, text)
	page.EndText()

	page.SetFontAndSize(font, fsize)
	page.SetRGBFill(c)
}

func printText(page goharu.Page, no *int) {
	pos := page.GetCurrentTextPos()
	*no = *no + 1
	buf := fmt.Sprintf(".[%d]%0.2f %0.2f", *no, pos.X, pos.Y)
	page.ShowText(buf)
}
