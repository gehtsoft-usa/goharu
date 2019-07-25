package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#cgo LDFLAGS: -lharu23
//#cgo windows,386 LDFLAGS: -L${SRCDIR}/libharu/windows/x86
//#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/libharu/windows/x64
//#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/libharu/macos/x64
//#cgo linux,!android,amd64 LDFLAGS: -L${SRCDIR}/libharu/linux/x64
//#cgo linux,!android,arm32 LDFLAGS: -L${SRCDIR}/libharu/linux/arm32
//#cgo linux,!android,arm64 LDFLAGS: -L${SRCDIR}/libharu/linux/arm64
//#include <hpdf.h>
//#include <stdio.h>
//extern void go_haru_error_callback(int error_no, int detail_no, int user_data);
//static void my_haru_error_callback(int error_no, int detail_no, int user_data)
//{
//		go_haru_error_callback(error_no, detail_no, user_data);
//}
//static HPDF_Doc my_create()
//{
//		return HPDF_New((HPDF_Error_Handler)my_haru_error_callback, (void*)0);
//}
import (
	"C"
)
import (
	"fmt"
	"unsafe"
)

func Version() string {
	return C.GoString(C.HPDF_GetVersion())
}

type Haru struct {
	pdf C.HPDF_Doc
}

//export go_haru_error_callback
func go_haru_error_callback(error_no, detail_no, user_data int32) {
	panic(fmt.Sprintf("Haru error %x/%x", error_no, detail_no))
}

func Create() Haru {
	var pdf C.HPDF_Doc
	pdf = C.my_create()
	return Haru{pdf: pdf}
}

func (v *Haru) NewDoc() {
	C.HPDF_NewDoc(v.pdf)
}

func (v *Haru) FreeDoc() {
	C.HPDF_FreeDoc(v.pdf)
}

func (v *Haru) Free() {
	C.HPDF_Free(v.pdf)
}

func (v *Haru) Content() []byte {
	var size C.uint
	var sizePtr *C.uint = &size
	C.HPDF_SaveToStream(v.pdf)
	size = C.HPDF_GetStreamSize(v.pdf)
	buff := C.malloc(C.size_t(size))
	defer C.free(buff)
	C.HPDF_GetContents(v.pdf, (*C.uchar)(buff), sizePtr)
	return C.GoBytes(buff, C.int(size))
}

func (v *Haru) Save(name string) {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.HPDF_SaveToFile(v.pdf, _name)
}
