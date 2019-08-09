package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//U3D struct
type U3D struct {
	ptr C.HPDF_U3D
}
