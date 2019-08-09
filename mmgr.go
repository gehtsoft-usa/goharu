package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//MMgr struct contains the reference to media manager
type MMgr struct {
	ptr C.HPDF_MMgr
}

//GetPageMMgr returns media manager object
func (page *Page) GetPageMMgr() MMgr {
	rc := C.HPDF_GetPageMMgr(page.ptr)
	return MMgr{ptr: rc}
}
