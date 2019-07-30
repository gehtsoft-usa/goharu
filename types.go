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

func (color *C.HPDF_RGBColor) toGo() RGBColor {
	return RGBColor{R: float32(color.r), G: float32(color.g), B: float32(color.b)}
}

type CMYKColor struct {
	C, M, Y, K float32
}

func (color *CMYKColor) toHaru() C.HPDF_CMYKColor {
	return C.HPDF_CMYKColor{c: C.float(color.C), m: C.float(color.M), y: C.float(color.Y), k: C.float(color.K)}
}

func (color *C.HPDF_CMYKColor) toGo() CMYKColor {
	return CMYKColor{C: float32(color.c), M: float32(color.m), Y: float32(color.y), K: float32(color.k)}
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

type TransMatrix struct {
	A, B, C, D, X, Y float32
}

func (src *C.HPDF_TransMatrix) toGo() TransMatrix {
	return TransMatrix{A: float32(src.a), B: float32(src.b), C: float32(src.c), D: float32(src.d), X: float32(src.x), Y: float32(src.y)}
}

func (src *TransMatrix) toHaru() C.HPDF_TransMatrix {
	return C.HPDF_TransMatrix{a: C.float(src.A), b: C.float(src.B), c: C.float(src.C), d: C.float(src.D), x: C.float(src.X), y: C.float(src.Y)}
}

type DashMode struct {
	Pattern     [8]uint16
	NumPatterns uint
	Phase       uint
}

func (src *DashMode) toHaru() C.HPDF_DashMode {
	r := C.HPDF_DashMode{num_ptn: C.uint(src.NumPatterns), phase: C.uint(src.Phase)}
	var i int
	for i = 0; i < 8; i++ {
		r.ptn[i] = C.ushort(src.Pattern[i])
	}
	return r
}

func (src *C.HPDF_DashMode) toGo() DashMode {
	r := DashMode{NumPatterns: uint(src.num_ptn), Phase: uint(src.phase)}
	var i uint
	for i = 0; i < r.NumPatterns; i++ {
		r.Pattern[i] = uint16(src.ptn[i])
	}
	return r
}
