package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//Outline struct keeps outline 
type Outline struct {
	ptr C.HPDF_Outline
}

//CreateOutline creates a new outline object.
func (pdf *Doc) CreateOutline(parent Outline, title string, encoder Encoder) Outline {
	_title := C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	rc := C.HPDF_CreateOutline(pdf.ptr, parent.ptr, _title, encoder.ptr)
	return Outline{ptr: rc}
}

//SetOpened sets flag indicating whether the outline is opened
func (outline *Outline) SetOpened(opened bool) {
	var _opened C.int
	if opened {
		_opened = C.HPDF_TRUE
	} else {
		_opened = C.HPDF_FALSE
	}

	C.HPDF_Outline_SetOpened(outline.ptr, _opened)
}

//SetDestination sets outline destination
func (outline *Outline) SetDestination(dst Destination) {

	C.HPDF_Outline_SetDestination(outline.ptr, dst.ptr)
}
