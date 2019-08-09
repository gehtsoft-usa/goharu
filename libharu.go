package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#cgo LDFLAGS: -lharu23
//#cgo windows,386 LDFLAGS: -L${SRCDIR}/libharu/windows/x86
//#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/libharu/windows/x64
//#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/libharu/macos/x64 -liconv
//#cgo linux,!android,amd64 LDFLAGS: -L${SRCDIR}/libharu/linux/x64 -lm
//#cgo linux,!android,arm32 LDFLAGS: -L${SRCDIR}/libharu/linux/arm32 -lm
//#cgo linux,!android,arm64 LDFLAGS: -L${SRCDIR}/libharu/linux/arm64 -lm
//#include <hpdf.h>
//#include <hpdf_u3d.h>
//extern void go_haru_error_callback(int error_no, int detail_no, int user_data);
//static void my_haru_error_callback(int error_no, int detail_no, int user_data)
//{
//      go_haru_error_callback(error_no, detail_no, user_data);
//}
//static HPDF_Doc my_create(int panicAtError)
//{
//      return HPDF_New(panicAtError != 0 ? (HPDF_Error_Handler)my_haru_error_callback : (void*)0, (void*)0);
//}
import (
	"C"
)

import (
	"fmt"
	"time"
	"unsafe"
)

//Version returns the version of the haru library
func Version() string {
	return C.GoString(C.HPDF_GetVersion())
}

//The Doc structure represents the haru engine
type Doc struct {
	ptr C.HPDF_Doc
}

//export go_haru_error_callback
func go_haru_error_callback(errorNo, detailNo, userData int32) {
	panic(fmt.Sprintf("Haru error %x/%x", errorNo, detailNo))
}

//Create initializes the haru engine
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

//NewDoc creates a new document.
func (pdf *Doc) NewDoc() {
	C.HPDF_NewDoc(pdf.ptr)
}

//FreeDoc revokes the current document.
func (pdf *Doc) FreeDoc() {
	C.HPDF_FreeDoc(pdf.ptr)
}

//Free revokes a document object and all resources.
func (pdf *Doc) Free() {
	C.HPDF_Free(pdf.ptr)
}

//Content gets the content of the document as a byte array
func (pdf *Doc) Content() []byte {
	var size C.uint
	var sizePtr *C.uint = &size
	C.HPDF_SaveToStream(pdf.ptr)
	size = C.HPDF_GetStreamSize(pdf.ptr)
	buff := C.malloc(C.size_t(size))
	defer C.free(buff)
	C.HPDF_GetContents(pdf.ptr, (*C.uchar)(buff), sizePtr)
	C.HPDF_ResetStream(pdf.ptr)
	return C.GoBytes(buff, C.int(size))
}

//Save saves the content of the document to the file specified
func (pdf *Doc) Save(name string) {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.HPDF_SaveToFile(pdf.ptr, _name)
}

//GetError gets last error code
//
//The function is used to check error if panicIfError parameter is not set in New function
func (pdf *Doc) GetError() uint32 {
	rc := C.HPDF_GetError(pdf.ptr)
	return uint32(rc)
}

//GetErrorDetail gets last error detail code
//
//The function is used to check error if panicIfError parameter is not set in New function
//TBD: GetErrorDetail
func (pdf *Doc) GetErrorDetail() uint32 {
	rc := C.HPDF_GetErrorDetail(pdf.ptr)
	return uint32(rc)
}

//ResetError resets error
//
//The function is used to reset error before continue if panicIfError parameter is not set in New function
func (pdf *Doc) ResetError() {
	C.HPDF_ResetError(pdf.ptr)
}

//InfoCreationDate is an identifier of a attribute
const InfoCreationDate int = 0

//InfoModDate is an identifier of a attribute
const InfoModDate int = 1

//SetInfoAttrDate  sets data/time attribute value
//
//Attribute code may be InfoCreationDate or InfoModDate
func (pdf *Doc) SetInfoAttrDate(attr int, t time.Time) {
	C.HPDF_SetInfoDateAttr(pdf.ptr, C.HPDF_InfoType(attr), timeToHaru(t))
}

//InfoAuthor is an identifier of an attribute
const InfoAuthor int = 2

//InfoCreator is an identifier of an attribute
const InfoCreator int = 3

//InfoProducer is an identifier of an attribute
const InfoProducer int = 4

//InfoTitle is an identifier of an attribute
const InfoTitle int = 5

//InfoSubject is an identifier of an attribute
const InfoSubject int = 6

//InfoKeywords is an identifier of an attribute
const InfoKeywords int = 7

//InfoTrapped is an identifier of an attribute
const InfoTrapped int = 8

//InfoGtsPdfX is an identifier of an attribute
const InfoGtsPdfX int = 9

//SetInfoAttrString sets string attribute value
//
//The attr value may be any of Info* except those with DATE
func (pdf *Doc) SetInfoAttrString(attr int, value string) {
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_value))
	C.HPDF_SetInfoAttr(pdf.ptr, C.HPDF_InfoType(attr), _value)
}

//GetInfoAttr gets info attribute value
//
//The attr value may be any of INFO_*
//GetInfoAttr gets an attribute value from info dictionary.
func (pdf *Doc) GetInfoAttr(attr int) string {
	return C.GoString(C.HPDF_GetInfoAttr(pdf.ptr, C.HPDF_InfoType(attr)))
}

//SetCompressionMode sets compression mode
//
//The mode may be any Comp* value except CompMask
//SetCompressionMode set the mode of compression.
func (pdf *Doc) SetCompressionMode(mode int) {
	C.HPDF_SetCompressionMode(pdf.ptr, C.uint(mode))
}

//SetPassword sets a password for the document.
func (pdf *Doc) SetPassword(ownerPassword string, userPassword string) {
	_ownerPassword := C.CString(ownerPassword)
	defer C.free(unsafe.Pointer(_ownerPassword))

	_userPassword := C.CString(userPassword)
	defer C.free(unsafe.Pointer(_userPassword))
	C.HPDF_SetPassword(pdf.ptr, _ownerPassword, _userPassword)
}

//SetPermission sets permission
//
//permission may be combination of Enable* flags
//SetPermission set the permission flags for the document.
func (pdf *Doc) SetPermission(permission uint) {
	C.HPDF_SetPermission(pdf.ptr, C.uint(permission))
}

//EncryptR2 is encryption mode
const EncryptR2 int = 2

//EncryptR3 is encryption mode
const EncryptR3 int = 3

//SetEncryptionMode set the encryption mode.
//
//use EncryptR2 or EncryptR3 as mode
func (pdf *Doc) SetEncryptionMode(mode int, keyLen uint) {
	C.HPDF_SetEncryptionMode(pdf.ptr, C.HPDF_EncryptMode(mode), C.uint(keyLen))
}

//SetPagesConfiguration sets count of pages per printed page
func (pdf *Doc) SetPagesConfiguration(pagePerPages uint) {
	C.HPDF_SetPagesConfiguration(pdf.ptr, C.uint(pagePerPages))
}

//SetOpenAction set the first page to appear when a document is opened.
func (pdf *Doc) SetOpenAction(openAction Destination) {
	C.HPDF_SetOpenAction(pdf.ptr, openAction.ptr)
}

//AttachFile attaches the file to the document
func (pdf *Doc) AttachFile(file string) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	C.HPDF_AttachFile(pdf.ptr, _file)
}

//OutputIntent struct
type OutputIntent struct {
	ptr C.HPDF_OutputIntent
}

//LoadIccProfileFromFile loads color profile from the file
func (pdf *Doc) LoadIccProfileFromFile(iccFileName string, numComponent int) OutputIntent {
	_iccFileName := C.CString(iccFileName)
	defer C.free(unsafe.Pointer(_iccFileName))
	return OutputIntent{ptr: C.HPDF_LoadIccProfileFromFile(pdf.ptr, _iccFileName, C.int(numComponent))}
}
