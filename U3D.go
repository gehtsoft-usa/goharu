package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//The structure represents an image
type U3D struct {
	ptr C.HPDF_U3D
}
