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
//#include <hpdf_u3d.h>
//extern void go_haru_error_callback(int error_no, int detail_no, int user_data);
//static void my_haru_error_callback(int error_no, int detail_no, int user_data)
//{
//		go_haru_error_callback(error_no, detail_no, user_data);
//}
//static HPDF_Doc my_create(int panicAtError)
//{
//		return HPDF_New(panicAtError != 0 ? (HPDF_Error_Handler)my_haru_error_callback : (void*)0, (void*)0);
//}
import (
	"C"
)

import (
	"fmt"
	"time"
	"unsafe"
)

//Returns the version of the haru library
func Version() string {
	return C.GoString(C.HPDF_GetVersion())
}

//The structure represents the haru engine
type Doc struct {
	ptr C.HPDF_Doc
}

//export go_haru_error_callback
func go_haru_error_callback(error_no, detail_no, user_data int32) {
	panic(fmt.Sprintf("Haru error %x/%x", error_no, detail_no))
}

//The function initializes the haru engine
func Create(panicAtError bool) Doc {
	var ptr C.HPDF_Doc
	var _panicAtError C.int
	if panicAtError {
		_panicAtError = 1
	} else {
		_panicAtError = 0
	}
	ptr = C.my_create(_panicAtError)
	return Doc{ptr: ptr}
}

//The function creates a new document
func (v *Doc) NewDoc() {
	C.HPDF_NewDoc(v.ptr)
}

//The function frees the document
func (v *Doc) FreeDoc() {
	C.HPDF_FreeDoc(v.ptr)
}

//The function frees the engine and all allocated resources
func (v *Doc) Free() {
	C.HPDF_Free(v.ptr)
}

//Gets the content of the document as a byte array
func (v *Doc) Content() []byte {
	var size C.uint
	var sizePtr *C.uint = &size
	C.HPDF_SaveToStream(v.ptr)
	size = C.HPDF_GetStreamSize(v.ptr)
	buff := C.malloc(C.size_t(size))
	defer C.free(buff)
	C.HPDF_GetContents(v.ptr, (*C.uchar)(buff), sizePtr)
	C.HPDF_ResetStream(v.ptr)
	return C.GoBytes(buff, C.int(size))
}

//Saves the content of the document to the file specified
func (v *Doc) Save(name string) {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.HPDF_SaveToFile(v.ptr, _name)
}

//Gets last error code
//
//The function is used to check error if panicIfError parameter is not set in New function
func (pdf *Doc) GetError() uint32 {
	rc := C.HPDF_GetError(pdf.ptr)
	return uint32(rc)
}

//Gets last error detail code
//
//The function is used to check error if panicIfError parameter is not set in New function
func (pdf *Doc) GetErrorDetail() uint32 {
	rc := C.HPDF_GetErrorDetail(pdf.ptr)
	return uint32(rc)
}

//Resets error
//
//The function is used to check error if panicIfError parameter is not set in New function
func (pdf *Doc) ResetError() {
	C.HPDF_ResetError(pdf.ptr)
}

const INFO_CREATION_DATE int = 0
const INFO_MOD_DATE int = 1

//Sets data/time attribute value
//
//Attribute code may be INFO_CREATION_DATE or INFO_MOD_DATE
func (v *Doc) SetInfoAttrDate(attr int, t time.Time) {
	C.HPDF_SetInfoDateAttr(v.ptr, C.HPDF_InfoType(attr), timeToHaru(t))
}

const INFO_AUTHOR int = 2
const INFO_CREATOR int = 3
const INFO_PRODUCER int = 4
const INFO_TITLE int = 5
const INFO_SUBJECT int = 6
const INFO_KEYWORDS int = 7
const INFO_TRAPPED int = 8
const INFO_GTS_PDF_X int = 9
const INFO_EOF int = 10

//Sets string attribute value
//
//The attr value may be any of INFO_* except those with DATE
func (v *Doc) SetInfoAttrString(attr int, value string) {
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_value))
	C.HPDF_SetInfoAttr(v.ptr, C.HPDF_InfoType(attr), _value)
}

//Gets info attribute value
//
//The attr value may be any of INFO_*
func (v *Doc) GetInfoAttr(attr int) string {
	return C.GoString(C.HPDF_GetInfoAttr(v.ptr, C.HPDF_InfoType(attr)))
}

//Sets compression mode
//
//The mode may be any COMP_* value except COMP_MASK
func (pdf *Doc) SetCompressionMode(mode uint) {
	C.HPDF_SetCompressionMode(pdf.ptr, C.uint(mode))
}
