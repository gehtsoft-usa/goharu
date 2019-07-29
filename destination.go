package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//The structure represents an image
type Destination struct {
	ptr C.HPDF_Destination
}

func (page *Page) Page_CreateDestination() Destination {
	rc := C.HPDF_Page_CreateDestination(page.ptr)
	return Destination{ptr: rc}
}

func (dst *Destination) Destination_SetXYZ(left float32, top float32, zoom float32) {
	C.HPDF_Destination_SetXYZ(dst.ptr, C.float(left), C.float(top), C.float(zoom))
}

func (dst *Destination) Destination_SetFit() {
	C.HPDF_Destination_SetFit(dst.ptr)
}

func (dst *Destination) Destination_SetFitH(top float32) {
	C.HPDF_Destination_SetFitH(dst.ptr, C.float(top))
}

func (dst *Destination) Destination_SetFitV(left float32) {
	C.HPDF_Destination_SetFitV(dst.ptr, C.float(left))
}

func (dst *Destination) Destination_SetFitR(left float32, bottom float32, right float32, top float32) {
	C.HPDF_Destination_SetFitR(dst.ptr, C.float(left), C.float(bottom), C.float(right), C.float(top))
}

func (dst *Destination) Destination_SetFitB() {
	C.HPDF_Destination_SetFitB(dst.ptr)
}

func (dst *Destination) Destination_SetFitBH(top float32) {
	C.HPDF_Destination_SetFitBH(dst.ptr, C.float(top))
}

func (dst *Destination) Destination_SetFitBV(left float32) {
	C.HPDF_Destination_SetFitBV(dst.ptr, C.float(left))
}
