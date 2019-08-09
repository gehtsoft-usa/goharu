package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//Font struct keeps font reference
type Font struct {
	ptr C.HPDF_Font
}

//GetFont gets font by the name and encoding
func (pdf *Doc) GetFont(fontName string, encodingName string) Font {
	cFontName := C.CString(fontName)
	defer C.free(unsafe.Pointer(cFontName))

	var cEncodingName *C.char
	var rc C.HPDF_Font
	if len(encodingName) > 0 {
		cEncodingName = C.CString(encodingName)
		defer C.free(unsafe.Pointer(cEncodingName))
		rc = C.HPDF_GetFont(pdf.ptr, cFontName, cEncodingName)
	} else {
		rc = C.HPDF_GetFont(pdf.ptr, cFontName, nil)
	}
	return Font{ptr: rc}
}

//LoadType1FontFromFile loads a Type1 font from an external file and registers it in the document object. (See "Fonts and Encodings")
func (pdf *Doc) LoadType1FontFromFile(afmFileName string, dataFileName string) string {

	cAfmFileName := C.CString(afmFileName)
	defer C.free(unsafe.Pointer(cAfmFileName))

	cDataFileName := C.CString(dataFileName)
	defer C.free(unsafe.Pointer(cDataFileName))
	rc := C.HPDF_LoadType1FontFromFile(pdf.ptr, cAfmFileName, cDataFileName)

	return C.GoString(rc)

}

//LoadTTFontFromFile loads a TrueType font from an external file and register it to a document object. (See "Fonts and Encodings")
func (pdf *Doc) LoadTTFontFromFile(filename string, embedding bool) string {

	cFileName := C.CString(filename)
	defer C.free(unsafe.Pointer(cFileName))
	var cEmbedding C.int
	if embedding {
		cEmbedding = C.HPDF_TRUE
	} else {
		cEmbedding = C.HPDF_FALSE
	}
	rc := C.HPDF_LoadTTFontFromFile(pdf.ptr, cFileName, cEmbedding)
	return C.GoString(rc)
}

//LoadTTFontFromFile2 loads a TrueType font from an TrueType collection file and register it to a document object. (See "Fonts and Encodings")
func (pdf *Doc) LoadTTFontFromFile2(filename string, index uint, embedding bool) string {

	cFileName := C.CString(filename)
	defer C.free(unsafe.Pointer(cFileName))

	var cEmbedding C.int
	if embedding {
		cEmbedding = C.HPDF_TRUE
	} else {
		cEmbedding = C.HPDF_FALSE
	}

	rc := C.HPDF_LoadTTFontFromFile2(pdf.ptr, cFileName, C.uint(index), cEmbedding)

	return C.GoString(rc)
}

//UseJPFonts enables Japanese fonts. After HPDFUseJPFonts() is invoked, an application can use the following Japanese fonts.
func (pdf *Doc) UseJPFonts() {

	C.HPDF_UseJPFonts(pdf.ptr)
}

//UseKRFonts enables Korean fonts. After HPDFUseKRFonts() is invoked, an application can use the following Korean fonts.
func (pdf *Doc) UseKRFonts() {

	C.HPDF_UseKRFonts(pdf.ptr)
}

//UseCNSFonts enables simplified Chinese fonts. After HPDFUseCNSFonts() is invoked, an application can use the following simplified Chinese fonts.
func (pdf *Doc) UseCNSFonts() {

	C.HPDF_UseCNSFonts(pdf.ptr)
}

//UseCNTFonts enables traditional Chinese fonts. After HPDFUseCNTFonts() is invoked, an application can use the following traditional Chinese fonts.
func (pdf *Doc) UseCNTFonts() {
	C.HPDF_UseCNTFonts(pdf.ptr)
}

//GetFontName gets the name of the font.
func (font *Font) GetFontName() string {

	rc := C.HPDF_Font_GetFontName(font.ptr)

	return C.GoString(rc)

}

//GetEncodingName gets the encoding name of the font.
func (font *Font) GetEncodingName() string {

	rc := C.HPDF_Font_GetEncodingName(font.ptr)

	return C.GoString(rc)

}

//GetUnicodeWidth gets the width of a Unicode character in a specific font.
func (font *Font) GetUnicodeWidth(code uint16) int {

	rc := C.HPDF_Font_GetUnicodeWidth(font.ptr, C.ushort(code))
	return int(rc)

}

//GetBBox gets the bounding box of the font.
func (font *Font) GetBBox() Rect {

	rc := C.HPDF_Font_GetBBox(font.ptr)

	return Rect{Left: float32(rc.left), Top: float32(rc.top), Right: float32(rc.right), Bottom: float32(rc.bottom)}

}

//GetAscent gets the vertical ascent of the font.
func (font *Font) GetAscent() int {

	rc := C.HPDF_Font_GetAscent(font.ptr)

	return int(rc)

}

//GetDescent gets the vertical descent of the font.
func (font *Font) GetDescent() int {

	rc := C.HPDF_Font_GetDescent(font.ptr)

	return int(rc)

}

//GetXHeight gets the distance from the baseline of lowercase letters.
func (font *Font) GetXHeight() uint {

	rc := C.HPDF_Font_GetXHeight(font.ptr)

	return uint(rc)

}

//GetCapHeight gets the distance from the baseline of uppercase letters.
func (font *Font) GetCapHeight() uint {

	rc := C.HPDF_Font_GetCapHeight(font.ptr)

	return uint(rc)

}

//TextWidth gets total width of the text, number of characters, and number of words.
func (font *Font) TextWidth(text string, len uint) TextWidth {
	cText := (*C.uchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(cText))
	rc := C.HPDF_Font_TextWidth(font.ptr, cText, C.uint(len))
	return TextWidth{NumberOfCharacters: uint(rc.numchars), NumberOfWords: uint(rc.numwords), Width: uint(rc.width), NumSpace: uint(rc.numspace)}
}

//TextWidthEncoded gets total width of the text, number of characters, and number of words.
func (font *Font) TextWidthEncoded(text []byte, len uint) TextWidth {
	cText := (*C.uchar)(unsafe.Pointer(&text[0]))
	return font.textWidthEncoded(cText, len)
}

//TextWidthEncoded1 gets total width of the text, number of characters, and number of words.
func (font *Font) TextWidthEncoded1(text string, encoding string) (TextWidth, error) {
	v, l, e := CEncodedString(text, encoding)
	if e != nil {
		return TextWidth{}, e
	}
	defer C.free(unsafe.Pointer(v))
	return font.textWidthEncoded((*C.uchar)(unsafe.Pointer(v)), uint(l)), nil
}

//TBD: textWidthEncoded
func (font *Font) textWidthEncoded(cText *C.uchar, len uint) TextWidth {
	rc := C.HPDF_Font_TextWidth(font.ptr, cText, C.uint(len))
	return TextWidth{NumberOfCharacters: uint(rc.numchars), NumberOfWords: uint(rc.numwords), Width: uint(rc.width), NumSpace: uint(rc.numspace)}
}

//MeasureText calculates the byte length which can be included within the specified width.
func (font *Font) MeasureText(text string, len uint, width float32, fontsize float32, charspace float32, wordspace float32, wordwrap bool) (uint, float32) {
	var cWordWrap C.int

	cText := (*C.uchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(cText))

	if wordwrap {
		cWordWrap = C.HPDF_TRUE
	} else {
		cWordWrap = C.HPDF_FALSE
	}

	var realwidth C.float
	rc := C.HPDF_Font_MeasureText(font.ptr, cText, C.uint(len), C.float(width), C.float(fontsize), C.float(charspace), C.float(wordspace), cWordWrap, &realwidth)

	return uint(rc), float32(realwidth)
}

//MeasureTextEncoded calculates the byte length which can be included within the specified width.
func (font *Font) MeasureTextEncoded(text []byte, len uint, width float32, fontsize float32, charspace float32, wordspace float32, wordwrap bool) (uint, float32) {
	cText := (*C.uchar)(unsafe.Pointer(&text[0]))
	return font.measureTextEncoded(cText, len, width, fontsize, charspace, wordspace, wordwrap)
}

//MeasureTextEncoded1 calculates the byte length which can be included within the specified width.
func (font *Font) MeasureTextEncoded1(text string, encoding string, width float32, fontsize float32, charspace float32, wordspace float32, wordwrap bool) (uint, float32, error) {
	v, l, e := CEncodedString(text, encoding)
	if e != nil {
		return 0, 0, e
	}
	defer C.free(unsafe.Pointer(v))
	a, b := font.measureTextEncoded((*C.uchar)(unsafe.Pointer(v)), uint(l), width, fontsize, charspace, wordspace, wordwrap)
	return a, b, nil
}

//TBD: measureTextEncoded
func (font *Font) measureTextEncoded(cText *C.uchar, len uint, width float32, fontsize float32, charspace float32, wordspace float32, wordwrap bool) (uint, float32) {
	var cWordWrap C.int
	if wordwrap {
		cWordWrap = C.HPDF_TRUE
	} else {
		cWordWrap = C.HPDF_FALSE
	}

	var realwidth C.float
	rc := C.HPDF_Font_MeasureText(font.ptr, cText, C.uint(len), C.float(width), C.float(fontsize), C.float(charspace), C.float(wordspace), cWordWrap, &realwidth)

	return uint(rc), float32(realwidth)
}
