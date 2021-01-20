package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
//#include <string.h>
//void zeroTerminate(char *buff, size_t size, size_t maxLen) {
//  size_t i;
//  for (i = 0; i < 4 && i + size < maxLen; i++)
//      buff[size + i] = 0;
//}
import "C"
import "unsafe"
import "golang.org/x/text/encoding/ianaindex"

//Encoder is text encoder used by libharu
type Encoder struct {
    ptr C.HPDF_Encoder
}

//CEncodedString creates an ASCIZ native string in specified encoding
//
//The returned value must be freed using C.free
func CEncodedString(text string, encoding string) (*C.char, int, error) {

    if encoding == "UTF-8" || encoding == "UTF8" {
        return C.CString(text), 0, nil
    }

    var encodingTable, err = ianaindex.MIME.Encoding(encoding)

    if err != nil {
        return C.CString(text), 0, err
    }

    var encoder = encodingTable.NewEncoder()

    var s, err1 = encoder.String(text);

    if err1 != nil {
        return C.CString(text), 0, err1
    }

    return C.CString(s), len(s) + 1, nil
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
