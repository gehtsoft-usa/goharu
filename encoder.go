package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#cgo LDFLAGS: -liconv
//#include <hpdf.h>
//#include <string.h>
//#include <iconv.h>
//void zeroTerminate(char *buff, size_t size, size_t maxLen) {
//	size_t i;
//	for (i = 0; i < 4 && i + size < maxLen; i++)
//		buff[size + i] = 0;
//}
import "C"
import "unsafe"

//The structure represents an image
type Encoder struct {
	ptr C.HPDF_Encoder
}

func CEncodedString(text string, encoding string) (*C.char, int, error) {
	var src *C.char
	src = C.CString(text)
	if encoding == "UTF-8" || encoding == "UTF8" {
		return src, 0, nil
	}

	_srcEncoding := C.CString("UTF-8")
	defer C.free(unsafe.Pointer(_srcEncoding))
	_dstEncoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(_dstEncoding))

	icd, err := C.iconv_open(_dstEncoding, _srcEncoding)

	if icd == nil || err != nil {
		return (*C.char)(unsafe.Pointer(nil)), 0, err
	}
	defer C.iconv_close(icd)
	defer C.free(unsafe.Pointer(src))

	var srcLen C.size_t = C.strlen(src)
	if srcLen == 0 {
		var dst *C.char
		dst = (*C.char)(C.malloc(C.size_t(1)))
		C.zeroTerminate(dst, 0, 1)
		return dst, 0, nil
	}
	var temporaryLen C.size_t = C.size_t(srcLen * 4)
	var bytesLeft C.size_t = temporaryLen
	var resultLen C.size_t
	var temporary, temporary1 *C.char
	temporary = (*C.char)(C.malloc(C.size_t(temporaryLen)))
	temporary1 = temporary
	_, err = C.iconv(icd, &src, &srcLen, &temporary1, &bytesLeft)
	if err != nil {
		C.free(unsafe.Pointer(temporary))
		return (*C.char)(unsafe.Pointer(nil)), 0, err
	}
	resultLen = temporaryLen - bytesLeft
	C.zeroTerminate(temporary, resultLen, temporaryLen)
	return temporary, int(resultLen), nil
}

func BufferToASCIZ(text []byte) *C.char {
	size := C.size_t(len(text))
	size1 := C.size_t(len(text) + 4)
	out := C.malloc(size)
	_text := (*C.uchar)(unsafe.Pointer(&text[0]))
	C.memcpy(unsafe.Pointer(out), unsafe.Pointer(_text), size)
	C.zeroTerminate((*C.char)(out), size, size1)
	return (*C.char)(out)
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
