package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//Destination struct
type Destination struct {
	ptr C.HPDF_Destination
}

//CreateDestination creates a new destination object for the page.
func (page *Page) CreateDestination() Destination {
	rc := C.HPDF_Page_CreateDestination(page.ptr)
	return Destination{ptr: rc}
}

//SetXYZ sets destination coordinates
func (dst *Destination) SetXYZ(left float32, top float32, zoom float32) {
	C.HPDF_Destination_SetXYZ(dst.ptr, C.float(left), C.float(top), C.float(zoom))
}

//SetFit makes the destination to fit
func (dst *Destination) SetFit() {
	C.HPDF_Destination_SetFit(dst.ptr)
}

//SetFitH makes the destination to fit horizontally
func (dst *Destination) SetFitH(top float32) {
	C.HPDF_Destination_SetFitH(dst.ptr, C.float(top))
}

//SetFitV makes the destination to fit vertically
func (dst *Destination) SetFitV(left float32) {
	C.HPDF_Destination_SetFitV(dst.ptr, C.float(left))
}

//SetFitR set destination to fit a rectangle
func (dst *Destination) SetFitR(left float32, bottom float32, right float32, top float32) {
	C.HPDF_Destination_SetFitR(dst.ptr, C.float(left), C.float(bottom), C.float(right), C.float(top))
}

//SetFitB ???
func (dst *Destination) SetFitB() {
	C.HPDF_Destination_SetFitB(dst.ptr)
}

//SetFitBH ???
func (dst *Destination) SetFitBH(top float32) {
	C.HPDF_Destination_SetFitBH(dst.ptr, C.float(top))
}

//SetFitBV ???
func (dst *Destination) SetFitBV(left float32) {
	C.HPDF_Destination_SetFitBV(dst.ptr, C.float(left))
}
