package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//The structure represents an image
type MMgr struct {
	ptr C.HPDF_MMgr
}

func (page *Page) GetPageMMgr() MMgr {
	rc := C.HPDF_GetPageMMgr(page.ptr)
	return MMgr{ptr: rc}
}
