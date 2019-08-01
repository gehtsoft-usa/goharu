package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//The structure represents an image
type Outline struct {
	ptr C.HPDF_Outline
}

func (pdf *Doc) CreateOutline(parent Outline, title string, encoder Encoder) Outline {
	_title := C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	rc := C.HPDF_CreateOutline(pdf.ptr, parent.ptr, _title, encoder.ptr)
	return Outline{ptr: rc}
}

func (outline *Outline) SetOpened(opened bool) {
	var _opened C.int
	if opened {
		_opened = C.HPDF_TRUE
	} else {
		_opened = C.HPDF_FALSE
	}

	C.HPDF_Outline_SetOpened(outline.ptr, _opened)
}

func (outline *Outline) SetDestination(dst Destination) {

	C.HPDF_Outline_SetDestination(outline.ptr, dst.ptr)
}
