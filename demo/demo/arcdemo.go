package demo

import (
	"github.com/gehtsoft-usa/goharu"
)

func Arcdemo() {
	pdf := goharu.Create(true)
	defer pdf.Free()
	pdf.NewDoc()
	defer pdf.FreeDoc()

	page := pdf.AddPage()

	page.SetHeight(220)
	page.SetWidth(200)

	printGrid(pdf, page)

	/* draw pie chart
	 *
	 *   A: 45% Red
	 *   B: 25% Blue
	 *   C: 15% green
	 *   D: other yellow
	 */

	var pos goharu.Point

	/* A */
	page.SetRGBFill(goharu.RGBColor{R: 1.0, G: 0.0, B: 0.0})
	page.MoveTo(100, 100)
	page.LineTo(100, 180)
	page.Arc(100, 100, 80, 0, 360*0.45)
	pos = page.GetCurrentPos()
	page.LineTo(100, 100)
	page.Fill()

	/* B */
	page.SetRGBFill(goharu.RGBColor{R: 0.0, G: 0.0, B: 1.0})
	page.MoveTo(100, 100)
	page.LineTo(pos.X, pos.Y)
	page.Arc(100, 100, 80, 360*0.45, 360*0.7)
	pos = page.GetCurrentPos()
	page.LineTo(100, 100)
	page.Fill()

	/* C */
	page.SetRGBFill(goharu.RGBColor{R: 0.0, G: 1.0, B: 0.0})
	page.MoveTo(100, 100)
	page.LineTo(pos.X, pos.Y)
	page.Arc(100, 100, 80, 360*0.7, 360*0.85)
	pos = page.GetCurrentPos()
	page.LineTo(100, 100)
	page.Fill()

	/* D */
	page.SetRGBFill(goharu.RGBColor{R: 1.0, G: 1.0, B: 0.0})
	page.MoveTo(100, 100)
	page.LineTo(pos.X, pos.Y)
	page.Arc(100, 100, 80, 360*0.85, 360)
	pos = page.GetCurrentPos()
	page.LineTo(100, 100)
	page.Fill()

	/* draw center circle */
	page.SetGrayStroke(0)
	page.SetGrayFill(1)
	page.Circle(100, 100, 30)
	page.Fill()

	pdf.Save("out/arcdemo.pdf")
}
