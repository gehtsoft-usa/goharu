package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//The structure represents an image
type Font struct {
	ptr C.HPDF_Font
}

func (pdf *Doc) GetFont(font_name string, encoding_name string) Font {
	_font_name := C.CString(font_name)
	defer C.free(unsafe.Pointer(_font_name))

	_encoding_name := C.CString(encoding_name)
	defer C.free(unsafe.Pointer(_encoding_name))
	rc := C.HPDF_GetFont(pdf.ptr, _font_name, _encoding_name)

	return Font{ptr: rc}
}

func (pdf *Doc) LoadType1FontFromFile(afm_file_name string, data_file_name string) string {

	_afm_file_name := C.CString(afm_file_name)
	defer C.free(unsafe.Pointer(_afm_file_name))

	_data_file_name := C.CString(data_file_name)
	defer C.free(unsafe.Pointer(_data_file_name))
	rc := C.HPDF_LoadType1FontFromFile(pdf.ptr, _afm_file_name, _data_file_name)

	return C.GoString(rc)

}

func (pdf *Doc) LoadTTFontFromFile(file_name string, embedding bool) string {

	_file_name := C.CString(file_name)
	defer C.free(unsafe.Pointer(_file_name))
	var _embedding C.int
	if embedding {
		_embedding = C.HPDF_TRUE
	} else {
		_embedding = C.HPDF_FALSE
	}
	rc := C.HPDF_LoadTTFontFromFile(pdf.ptr, _file_name, _embedding)
	return C.GoString(rc)
}

func (pdf *Doc) LoadTTFontFromFile2(file_name string, index uint, embedding bool) string {

	_file_name := C.CString(file_name)
	defer C.free(unsafe.Pointer(_file_name))

	var _embedding C.int
	if embedding {
		_embedding = C.HPDF_TRUE
	} else {
		_embedding = C.HPDF_FALSE
	}

	rc := C.HPDF_LoadTTFontFromFile2(pdf.ptr, _file_name, C.uint(index), _embedding)

	return C.GoString(rc)
}

func (pdf *Doc) UseKRFonts() {

	C.HPDF_UseKRFonts(pdf.ptr)
}

func (pdf *Doc) UseCNSFonts() {

	C.HPDF_UseCNSFonts(pdf.ptr)
}

func (pdf *Doc) UseCNTFonts() {
	C.HPDF_UseCNTFonts(pdf.ptr)
}

func (font *Font) GetFontName() string {

	rc := C.HPDF_Font_GetFontName(font.ptr)

	return C.GoString(rc)

}

func (font *Font) GetEncodingName() string {

	rc := C.HPDF_Font_GetEncodingName(font.ptr)

	return C.GoString(rc)

}

func (font *Font) GetUnicodeWidth(code uint16) int {

	rc := C.HPDF_Font_GetUnicodeWidth(font.ptr, C.ushort(code))
	return int(rc)

}

func (font *Font) GetBBox() Rect {

	rc := C.HPDF_Font_GetBBox(font.ptr)

	return Rect{Left: float32(rc.left), Top: float32(rc.top), Right: float32(rc.right), Bottom: float32(rc.bottom)}

}

func (font *Font) GetAscent() int {

	rc := C.HPDF_Font_GetAscent(font.ptr)

	return int(rc)

}

func (font *Font) GetDescent() int {

	rc := C.HPDF_Font_GetDescent(font.ptr)

	return int(rc)

}

func (font *Font) GetXHeight() uint {

	rc := C.HPDF_Font_GetXHeight(font.ptr)

	return uint(rc)

}

func (font *Font) GetCapHeight() uint {

	rc := C.HPDF_Font_GetCapHeight(font.ptr)

	return uint(rc)

}

func (font *Font) TextWidth(text string, len uint) TextWidth {

	_text := (*C.uchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_text))

	rc := C.HPDF_Font_TextWidth(font.ptr, _text, C.uint(len))
	return TextWidth{NumberOfCharacters: uint(rc.numchars), NumberOfWords: uint(rc.numwords), Width: uint(rc.width), NumSpace: uint(rc.numspace)}

}

func (font *Font) TextWidthEncoded(text []byte, len uint) TextWidth {

	_text := (*C.uchar)(unsafe.Pointer(&text[0]))
	rc := C.HPDF_Font_TextWidth(font.ptr, _text, C.uint(len))
	return TextWidth{NumberOfCharacters: uint(rc.numchars), NumberOfWords: uint(rc.numwords), Width: uint(rc.width), NumSpace: uint(rc.numspace)}
}

func (font *Font) MeasureText(text string, len uint, width float32, font_size float32, char_space float32, word_space float32, wordwrap bool) (uint, float32) {
	var _wordwrap C.int

	_text := (*C.uchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_text))

	if wordwrap {
		_wordwrap = C.HPDF_TRUE
	} else {
		_wordwrap = C.HPDF_FALSE
	}

	var realwidth C.float
	rc := C.HPDF_Font_MeasureText(font.ptr, _text, C.uint(len), C.float(width), C.float(font_size), C.float(char_space), C.float(word_space), _wordwrap, &realwidth)

	return uint(rc), float32(realwidth)
}

func (font *Font) MeasureTextEncoded(text []byte, len uint, width float32, font_size float32, char_space float32, word_space float32, wordwrap bool) (uint, float32) {
	var _wordwrap C.int
	if wordwrap {
		_wordwrap = C.HPDF_TRUE
	} else {
		_wordwrap = C.HPDF_FALSE
	}
	_text := (*C.uchar)(unsafe.Pointer(&text[0]))

	var realwidth C.float
	rc := C.HPDF_Font_MeasureText(font.ptr, _text, C.uint(len), C.float(width), C.float(font_size), C.float(char_space), C.float(word_space), _wordwrap, &realwidth)

	return uint(rc), float32(realwidth)
}
