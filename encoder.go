package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//The structure represents an image
type Encoder struct {
	ptr C.HPDF_Encoder
}

func (pdf *Doc) GetEncoder(encoding_name string) Encoder {
	_encoding_name := C.CString(encoding_name)
	defer C.free(unsafe.Pointer(_encoding_name))
	rc := C.HPDF_GetEncoder(pdf.ptr, _encoding_name)
	return Encoder{ptr: rc}
}

func (pdf *Doc) GetCurrentEncoder() Encoder {
	rc := C.HPDF_GetCurrentEncoder(pdf.ptr)
	return Encoder{ptr: rc}
}

func (pdf *Doc) SetCurrentEncoder(encoding_name string) {
	_encoding_name := C.CString(encoding_name)
	defer C.free(unsafe.Pointer(_encoding_name))
	C.HPDF_SetCurrentEncoder(pdf.ptr, _encoding_name)
}

func (encoder *Encoder) Encoder_GetType() int {
	return int(C.HPDF_Encoder_GetType(encoder.ptr))
}

func (encoder *Encoder) Encoder_GetByteType(text string, index uint) int {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	return int(C.HPDF_Encoder_GetByteType(encoder.ptr, _text, C.uint(index)))
}

func (encoder *Encoder) Encoder_GetUnicode(code uint16) uint16 {
	return uint16(C.HPDF_Encoder_GetUnicode(encoder.ptr, C.ushort(code)))
}

const WRITING_MODE_HORIZONTAL int = 0
const WRITING_MODE_VERTICAL int = 1

func (encoder *Encoder) Encoder_GetWritingMode() int {
	return int(C.HPDF_Encoder_GetWritingMode(encoder.ptr))
}

func (pdf *Doc) UseJPEncodings() {
	C.HPDF_UseJPEncodings(pdf.ptr)
}

func (pdf *Doc) UseKREncodings() {
	C.HPDF_UseKREncodings(pdf.ptr)
}

func (pdf *Doc) UseCNSEncodings() {
	C.HPDF_UseCNSEncodings(pdf.ptr)
}

func (pdf *Doc) UseCNTEncodings() {
	C.HPDF_UseCNTEncodings(pdf.ptr)
}

func (pdf *Doc) UseUTFEncodings() {
	C.HPDF_UseUTFEncodings(pdf.ptr)
}
