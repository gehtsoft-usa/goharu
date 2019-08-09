package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//ExData struct
type ExData struct {
	ptr C.HPDF_ExData
}
