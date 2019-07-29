package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
//#include <hpdf_u3d.h>
import "C"
import "unsafe"

//The structure represents an image
type JavaScript struct {
	ptr C.HPDF_Dict
}

func (v *Doc) LoadJavaScriptFromFile(filename string) JavaScript {
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	ptr := C.HPDF_LoadJSFromFile(v.ptr, _filename)
	return JavaScript{ptr: ptr}
}

func (v *Doc) CreateJavaScript(code string) JavaScript {
	_code := C.CString(code)
	defer C.free(unsafe.Pointer(_code))
	ptr := C.HPDF_CreateJavaScript(v.ptr, _code)
	return JavaScript{ptr: ptr}
}
