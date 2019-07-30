package demo

import (
	"fmt"
	"math"

	"github.com/gehtsoft-usa/goharu"
)

func printGrid(doc goharu.Doc, page goharu.Page) {
	{
		height := page.GetHeight()
		width := page.GetWidth()
		font := doc.GetFont("Helvetica", "")
		var x, y float32

		page.SetFontAndSize(font, 5)
		page.SetGrayFill(0.5)
		page.SetGrayStroke(0.8)

		/* Draw horizontal lines */
		y = 0
		for y < height {
			if math.Mod(float64(y), 10) == 0 {
				page.SetLineWidth(0.5)
			} else {
				if page.GetLineWidth() != 0.25 {
					page.SetLineWidth(0.25)
				}
			}

			page.MoveTo(0, y)
			page.LineTo(width, y)
			page.Stroke()

			if math.Mod(float64(y), 10) == 0 && y > 0 {
				page.SetGrayStroke(0.5)

				page.MoveTo(0, y)
				page.LineTo(5, y)
				page.Stroke()

				page.SetGrayStroke(0.8)
			}

			y += 5
		}

		/* Draw vertical lines */
		x = 0
		for x < width {
			if math.Mod(float64(x), 10) == 0 {
				page.SetLineWidth(0.5)
			} else {
				if page.GetLineWidth() != 0.25 {
					page.SetLineWidth(0.25)
				}
			}

			page.MoveTo(x, 0)
			page.LineTo(x, height)
			page.Stroke()

			if math.Mod(float64(x), 50) == 0 && x > 0 {
				page.SetGrayStroke(0.5)

				page.MoveTo(x, 0)
				page.LineTo(x, 5)
				page.Stroke()

				page.MoveTo(x, height)
				page.LineTo(x, height-5)
				page.Stroke()

				page.SetGrayStroke(0.8)
			}

			x += 5
		}

		/* Draw horizontal text */
		y = 0
		for y < height {
			if math.Mod(float64(y), 10) == 0 && y > 0 {
				page.BeginText()
				page.MoveTextPos(5, y-2)
				page.ShowText(fmt.Sprintf("%.0f", y))
				page.EndText()
			}
			y += 5
		}

		/* Draw vertical text */
		x = 0
		for x < width {
			if math.Mod(float64(x), 50) == 0 && x > 0 {
				s := fmt.Sprintf("%.0f", x)
				page.BeginText()
				page.MoveTextPos(x, 5)
				page.ShowText(s)
				page.EndText()

				page.BeginText()
				page.MoveTextPos(x, height-10)
				page.ShowText(s)
				page.EndText()
			}

			x += 5
		}

		page.SetGrayFill(0)
		page.SetGrayStroke(0)
	}
}
