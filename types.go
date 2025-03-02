package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "time"

//RGBColor keeps a color as RGB quad.
//
//The components are expressed as value between 0 and 1.
type RGBColor struct {
	R, G, B float32
}

//TBD: toHaru
func (color *RGBColor) toHaru() C.HPDF_RGBColor {
	return C.HPDF_RGBColor{r: C.float(color.R), g: C.float(color.G), b: C.float(color.B)}
}

func RGBColorFromNative(color C.HPDF_RGBColor) RGBColor {
	return RGBColor{R: float32(color.r), G: float32(color.g), B: float32(color.b)}
}

//func (native *C.HPDF_RGBColor) toGo() RGBColor {
//	return RGBColor{R: float32(native.r), G: float32(native.g), B: float32(native.b)}
//}

//func (native *_Ctype_struct__HPDF_RGBColor) toGo() RGBColor {
//	return RGBColor{R: float32(native.r), G: float32(native.g), B: float32(native.b)}
//}

//CMYKColor keeps a color as CMYK color.
//
//The components are expressed as value between 0 and 1.
type CMYKColor struct {
	C, M, Y, K float32
}

//TBD: toHaru
func (color *CMYKColor) toHaru() C.HPDF_CMYKColor {
	return C.HPDF_CMYKColor{c: C.float(color.C), m: C.float(color.M), y: C.float(color.Y), k: C.float(color.K)}
}

func CMYKColorFromNative(color C.HPDF_CMYKColor) CMYKColor {
	return CMYKColor{C: float32(color.c), M: float32(color.m), Y: float32(color.y), K: float32(color.k)}
}

//func (native *C.HPDF_CMYKColor) toGo() CMYKColor {
//	return CMYKColor{C: float32(native.c), M: float32(native.m), Y: float32(native.y), K: float32(native.k)}
//}

//func (native *_Ctype_struct__HPDF_CMYKColor) toGo() CMYKColor {
//	return CMYKColor{C: float32(native.c), M: float32(native.m), Y: float32(native.y), K: float32(native.k)}
//}

//Rect struct is used to store a rectangle coordinates
type Rect struct {
	Top, Bottom, Left, Right float32
}

//TBD: toHaru
func (rect *Rect) toHaru() C.HPDF_Rect {
	return C.HPDF_Rect{left: C.float(rect.Left), top: C.float(rect.Top), right: C.float(rect.Right), bottom: C.float(rect.Bottom)}
}

//Point struct is used to store 2D point coordinates
type Point struct {
	X, Y float32
}

//TBD: toHaru
func (point *Point) toHaru() C.HPDF_Point {
	return C.HPDF_Point{x: C.float(point.X), y: C.float(point.Y)}
}

//Point3D struct is used to store 3D point coordinates
type Point3D struct {
	X, Y, Z float32
}

//TBD: toHaru
func (point *Point3D) toHaru() C.HPDF_Point3D {
	return C.HPDF_Point3D{x: C.float(point.X), y: C.float(point.Y), z: C.float(point.Z)}
}

//TextWidth struct keeps complex results returned
//by text measurement methods
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

//TransMatrix stores information about at transformation matrix
type TransMatrix struct {
	A, B, C, D, X, Y float32
}

func TransMatrixFromNative(native C.HPDF_TransMatrix) TransMatrix {
	return TransMatrix{A: float32(native.a), B: float32(native.b), C: float32(native.c), D: float32(native.d), X: float32(native.x), Y: float32(native.y)}
}

//func (native *_Ctype_struct__HPDF_TransMatrix) toGo() TransMatrix {
//	return TransMatrix{A: float32(native.a), B: float32(native.b), C: float32(native.c), D: float32(native.d), X: float32(native.x), Y: float32(native.y)}
//}

//TBD: toHaru
func (tm *TransMatrix) toHaru() C.HPDF_TransMatrix {
	return C.HPDF_TransMatrix{a: C.float(tm.A), b: C.float(tm.B), c: C.float(tm.C), d: C.float(tm.D), x: C.float(tm.X), y: C.float(tm.Y)}
}

//DashMode struct is used to store information on how to draw one dash of a dashed line
type DashMode struct {
	Pattern     [8]uint16
	NumPatterns uint
	Phase       uint
}

//TBD: toHaru
func (src *DashMode) toHaru() C.HPDF_DashMode {
	r := C.HPDF_DashMode{num_ptn: C.uint(src.NumPatterns), phase: C.uint(src.Phase)}
	var i int
	for i = 0; i < 8; i++ {
		r.ptn[i] = C.ushort(src.Pattern[i])
	}
	return r
}

func DashModeFromNative(native C.HPDF_DashMode) DashMode {
	r := DashMode{NumPatterns: uint(native.num_ptn), Phase: uint(native.phase)}
	var i uint
	for i = 0; i < r.NumPatterns; i++ {
		r.Pattern[i] = uint16(native.ptn[i])
	}
	return r
}

//func (native *_Ctype_struct__HPDF_DashMode) toGo() DashMode {
//	r := DashMode{NumPatterns: uint(native.num_ptn), Phase: uint(native.phase)}
//	var i uint
//	for i = 0; i < r.NumPatterns; i++ {
//		r.Pattern[i] = uint16(native.ptn[i])
//	}
//	return r
//}
