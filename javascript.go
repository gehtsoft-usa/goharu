package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
//#include <hpdf_u3d.h>
import "C"
import "unsafe"

//JavaScript structure keeps information about a JavaScript
type JavaScript struct {
	ptr C.HPDF_Dict
}

//LoadJavaScriptFromFile loads JavaScript from a file
func (v *Doc) LoadJavaScriptFromFile(filename string) JavaScript {
	_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(_filename))
	ptr := C.HPDF_LoadJSFromFile(v.ptr, _filename)
	return JavaScript{ptr: ptr}
}

//CreateJavaScript creates JavaScript using the source code
func (v *Doc) CreateJavaScript(code string) JavaScript {
	_code := C.CString(code)
	defer C.free(unsafe.Pointer(_code))
	ptr := C.HPDF_CreateJavaScript(v.ptr, _code)
	return JavaScript{ptr: ptr}
}
