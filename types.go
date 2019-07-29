package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "time"

type RGBColor struct {
	R, G, B float32
}

func (color *RGBColor) toHaru() C.HPDF_RGBColor {
	return C.HPDF_RGBColor{r: C.float(color.R), g: C.float(color.G), b: C.float(color.B)}
}

type CMYKColor struct {
	C, M, Y, K float32
}

func (color *CMYKColor) toHaru() C.HPDF_CMYKColor {
	return C.HPDF_CMYKColor{c: C.float(color.C), m: C.float(color.M), y: C.float(color.Y), k: C.float(color.K)}
}

type Rect struct {
	Top, Bottom, Left, Right float32
}

func (rect *Rect) toHaru() C.HPDF_Rect {
	return C.HPDF_Rect{left: C.float(rect.Left), top: C.float(rect.Top), right: C.float(rect.Right), bottom: C.float(rect.Bottom)}
}

type Point struct {
	X, Y float32
}

func (point *Point) toHaru() C.HPDF_Point {
	return C.HPDF_Point{x: C.float(point.X), y: C.float(point.Y)}
}

type Point3D struct {
	X, Y, Z float32
}

func (point *Point3D) toHaru() C.HPDF_Point3D {
	return C.HPDF_Point3D{x: C.float(point.X), y: C.float(point.Y), z: C.float(point.Z)}
}

type TextWidth struct {
	NumberOfCharacters, NumberOfWords, Width, NumSpace uint
}

func timeToHaru(t time.Time) C.struct__HPDF_Date {
	var ht C.struct__HPDF_Date
	ht = C.struct__HPDF_Date{
		year:    C.int(t.Year()),
		month:   C.int(t.Month()),
		day:     C.int(t.Day()),
		hour:    C.int(t.Hour()),
		minutes: C.int(t.Minute()),
	}
	return ht
}
