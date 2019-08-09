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

//Encoder is text encoder used by libharu
type Encoder struct {
	ptr C.HPDF_Encoder
}

//CEncodedString creates an ASCIZ native string in specified encoding
//
//The returned value must be freed using C.free
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

//BufferToASCIZ converts a byte buffer into zero-terminated ASCIZ string
//
//The returned value must be freed using C.free
func BufferToASCIZ(text []byte) *C.char {
	size := C.size_t(len(text))
	size1 := C.size_t(len(text) + 4)
	out := C.malloc(size)
	_text := (*C.uchar)(unsafe.Pointer(&text[0]))
	C.memcpy(unsafe.Pointer(out), unsafe.Pointer(_text), size)
	C.zeroTerminate((*C.char)(out), size, size1)
	return (*C.char)(out)
}

//GetEncoder gets the handle of an encoder object by specified encoding name.
func (pdf *Doc) GetEncoder(encodingName string) Encoder {
	cEncodingName := C.CString(encodingName)
	defer C.free(unsafe.Pointer(cEncodingName))
	rc := C.HPDF_GetEncoder(pdf.ptr, cEncodingName)
	return Encoder{ptr: rc}
}

//GetCurrentEncoder gets the handle of the current encoder of the document object.
func (pdf *Doc) GetCurrentEncoder() Encoder {
	rc := C.HPDF_GetCurrentEncoder(pdf.ptr)
	return Encoder{ptr: rc}
}

//SetCurrentEncoder sets the current encoder for the document.
func (pdf *Doc) SetCurrentEncoder(encodingName string) {
	cEncodingName := C.CString(encodingName)
	defer C.free(unsafe.Pointer(cEncodingName))
	C.HPDF_SetCurrentEncoder(pdf.ptr, cEncodingName)
}

//GetType gets the type of an encoding object.
func (encoder *Encoder) GetType() int {
	return int(C.HPDF_Encoder_GetType(encoder.ptr))
}

//GetByteType returns the type of byte in the text at position index.
func (encoder *Encoder) GetByteType(text string, index uint) int {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	return int(C.HPDF_Encoder_GetByteType(encoder.ptr, _text, C.uint(index)))
}

//GetUnicode converts a specified character code to unicode.
func (encoder *Encoder) GetUnicode(code uint16) uint16 {
	return uint16(C.HPDF_Encoder_GetUnicode(encoder.ptr, C.ushort(code)))
}

//WritingModeHorizontal is encoding writing mode
const WritingModeHorizontal int = 0 
//WritingModeVertical is encoding writing mode
const WritingModeVertical int = 1 

//GetWritingMode returns the writing mode for the encoding object.
//
//Returns WritingMode* value 
func (encoder *Encoder) GetWritingMode() int {
	return int(C.HPDF_Encoder_GetWritingMode(encoder.ptr))
}

//UseJPEncodings enables Japanese encodings. After HPDF_UseJPEncodings() is invoked, an application can use the following Japanese encodings.
func (pdf *Doc) UseJPEncodings() {
	C.HPDF_UseJPEncodings(pdf.ptr)
}

//UseKREncodings enables Korean encodings. After HPDF_UseKREncodings() is invoked, an application can use the following Korean encodings.
func (pdf *Doc) UseKREncodings() {
	C.HPDF_UseKREncodings(pdf.ptr)
}

//UseCNSEncodings enables simplified Chinese encodings. After HPDF_UseCNSEncodings() is invoked, an application can use the following simplified Chinese encodings.
func (pdf *Doc) UseCNSEncodings() {
	C.HPDF_UseCNSEncodings(pdf.ptr)
}

//UseCNTEncodings enables traditional Chinese encodings. After HPDF_UseCNTEncodings() is invoked, an application can use the following traditional Chinese encodings.
func (pdf *Doc) UseCNTEncodings() {
	C.HPDF_UseCNTEncodings(pdf.ptr)
}

//UseUTFEncodings enables UTF-8 encodings. After HPDF_UseUTFEncodings() is invoked, an application can include UTF-8 encoded Unicode text (up to 3-byte UTF-8 sequences only). An application can use the following Unicode encodings (but only with TrueType fonts):
func (pdf *Doc) UseUTFEncodings() {
	C.HPDF_UseUTFEncodings(pdf.ptr)
}
