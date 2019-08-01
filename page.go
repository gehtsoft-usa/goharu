package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//The structure represents an image
type Page struct {
	ptr C.HPDF_Page
}

//Add a new page
func (pdf *Doc) AddPage() Page {
	rc := C.HPDF_AddPage(pdf.ptr)
	return Page{ptr: rc}
}

func (pdf *Doc) InsertPage(page Page) Page {
	rc := C.HPDF_InsertPage(pdf.ptr, page.ptr)
	return Page{ptr: rc}
}

func (pdf *Doc) GetCurrentPage() Page {
	rc := C.HPDF_GetCurrentPage(pdf.ptr)
	return Page{ptr: rc}
}

func (pdf *Doc) GetPageByIndex(index uint) Page {
	rc := C.HPDF_GetPageByIndex(pdf.ptr, C.uint(index))
	return Page{ptr: rc}
}

func (page *Page) SetWidth(value float32) {
	C.HPDF_Page_SetWidth(page.ptr, C.float(value))
}

func (page *Page) SetHeight(value float32) {
	C.HPDF_Page_SetHeight(page.ptr, C.float(value))
}

func (page *Page) SetRotate(angle uint16) {
	C.HPDF_Page_SetRotate(page.ptr, C.ushort(angle))
}

func (page *Page) SetZoom(zoom float32) {
	C.HPDF_Page_SetZoom(page.ptr, C.float(zoom))
}

const PAGE_NUM_STYLE_DECIMAL int = 0
const PAGE_NUM_STYLE_UPPER_ROMAN int = 1
const PAGE_NUM_STYLE_LOWER_ROMAN int = 2
const PAGE_NUM_STYLE_UPPER_LETTERS int = 3
const PAGE_NUM_STYLE_LOWER_LETTERS int = 4
const PAGE_NUM_STYLE_EOF int = 5

func (pdf *Doc) AddPageLabel(page_num uint, style int, first_page uint, prefix string) {
	_prefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(_prefix))
	C.HPDF_AddPageLabel(pdf.ptr, C.uint(page_num), C.HPDF_PageNumStyle(style), C.uint(first_page), _prefix)
}

const PAGE_LAYOUT_SINGLE int = 0
const PAGE_LAYOUT_ONE_COLUMN int = 1
const PAGE_LAYOUT_TWO_COLUMN_LEFT int = 2
const PAGE_LAYOUT_TWO_COLUMN_RIGHT int = 3
const PAGE_LAYOUT_TWO_PAGE_LEFT int = 4
const PAGE_LAYOUT_TWO_PAGE_RIGHT int = 5

func (pdf *Doc) GetPageLayout() int {
	rc := C.HPDF_GetPageLayout(pdf.ptr)
	return int(rc)
}

func (pdf *Doc) SetPageLayout(layout int) {
	C.HPDF_SetPageLayout(pdf.ptr, C.HPDF_PageLayout(layout))
}

const PAGE_MODE_USE_NONE int = 0
const PAGE_MODE_USE_OUTLINE int = 1
const PAGE_MODE_USE_THUMBS int = 2
const PAGE_MODE_FULL_SCREEN int = 3

func (pdf *Doc) GetPageMode() int {
	rc := C.HPDF_GetPageMode(pdf.ptr)
	return int(rc)
}

func (pdf *Doc) SetPageMode(mode int) {
	C.HPDF_SetPageMode(pdf.ptr, C.HPDF_PageMode(mode))
}

func (page *Page) TextWidth(text string) float32 {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	return float32(C.HPDF_Page_TextWidth(page.ptr, _text))
}

func (page *Page) MeasureText(text string, width float32, wordwrap bool) (uint, float32) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	var _wordwrap C.int
	if wordwrap {
		_wordwrap = C.HPDF_TRUE
	} else {
		_wordwrap = C.HPDF_FALSE
	}
	var _real_width C.float = 0

	rc := uint(C.HPDF_Page_MeasureText(page.ptr, _text, C.float(width), _wordwrap, &_real_width))
	return rc, float32(_real_width)
}

func (page *Page) GetWidth() float32 {
	return float32(C.HPDF_Page_GetWidth(page.ptr))
}

func (page *Page) GetHeight() float32 {
	return float32(C.HPDF_Page_GetHeight(page.ptr))
}

func (page *Page) GetGMode() uint16 {
	return uint16(C.HPDF_Page_GetGMode(page.ptr))
}

func (page *Page) GetCurrentPos() Point {
	rc := C.HPDF_Page_GetCurrentPos(page.ptr)
	return Point{X: float32(rc.x), Y: float32(rc.y)}

}

func (page *Page) GetCurrentTextPos() Point {
	rc := C.HPDF_Page_GetCurrentTextPos(page.ptr)
	return Point{X: float32(rc.x), Y: float32(rc.y)}
}

func (page *Page) GetCurrentFont() Font {
	rc := C.HPDF_Page_GetCurrentFont(page.ptr)
	return Font{ptr: rc}
}

func (page *Page) GetCurrentFontSize() float32 {
	return float32(C.HPDF_Page_GetCurrentFontSize(page.ptr))
}

func (page *Page) GetTransMatrix() TransMatrix {
	rc := C.HPDF_Page_GetTransMatrix(page.ptr)
	return rc.toGo()
}

func (page *Page) GetLineWidth() float32 {
	return float32(C.HPDF_Page_GetLineWidth(page.ptr))
}

const BUTT_END int = 0
const ROUND_END int = 1
const PROJECTING_SQUARE_END int = 2

const MITER_JOIN int = 0
const ROUND_JOIN int = 1
const BEVEL_JOIN int = 2

func (page *Page) GetLineCap() int {
	return int(C.HPDF_Page_GetLineCap(page.ptr))
}

func (page *Page) GetLineJoin() int {
	return int(C.HPDF_Page_GetLineJoin(page.ptr))
}

func (page *Page) GetMiterLimit() float32 {
	return float32(C.HPDF_Page_GetMiterLimit(page.ptr))
}

func (page *Page) GetDash() DashMode {
	rc := C.HPDF_Page_GetDash(page.ptr)
	return rc.toGo()
}

func (page *Page) GetFlat() float32 {
	return float32(C.HPDF_Page_GetFlat(page.ptr))
}

func (page *Page) GetCharSpace() float32 {
	return float32(C.HPDF_Page_GetCharSpace(page.ptr))
}

func (page *Page) GetWordSpace() float32 {
	return float32(C.HPDF_Page_GetWordSpace(page.ptr))
}

func (page *Page) GetHorizontalScalling() float32 {
	return float32(C.HPDF_Page_GetHorizontalScalling(page.ptr))
}

func (page *Page) GetTextLeading() float32 {
	return float32(C.HPDF_Page_GetTextLeading(page.ptr))
}

const TEXT_RENDER_FILL int = 0
const TEXT_RENDER_STROKE int = 1
const TEXT_RENDER_FILL_THEN_STROKE int = 2
const TEXT_RENDER_INVISIBLE int = 3
const TEXT_RENDER_FILL_CLIPPING int = 4
const TEXT_RENDER_STROKE_CLIPPING int = 5
const TEXT_RENDER_FILL_STROKE_CLIPPING int = 6
const TEXT_RENDER_CLIPPING int = 7

func (page *Page) GetTextRenderingMode() int {
	return int(C.HPDF_Page_GetTextRenderingMode(page.ptr))
}

func (page *Page) GetTextRaise() float32 {
	return float32(C.HPDF_Page_GetTextRaise(page.ptr))
}

func (page *Page) GetTextRise() float32 {
	return float32(C.HPDF_Page_GetTextRise(page.ptr))
}

func (page *Page) GetRGBFill() RGBColor {
	rc := C.HPDF_Page_GetRGBFill(page.ptr)
	return rc.toGo()
}

func (page *Page) GetRGBStroke() RGBColor {
	rc := C.HPDF_Page_GetRGBStroke(page.ptr)
	return rc.toGo()
}

func (page *Page) GetCMYKFill() CMYKColor {
	rc := C.HPDF_Page_GetCMYKFill(page.ptr)
	return rc.toGo()
}

func (page *Page) GetCMYKStroke() CMYKColor {
	rc := C.HPDF_Page_GetCMYKStroke(page.ptr)
	return rc.toGo()

}

func (page *Page) GetGrayFill() float32 {
	return float32(C.HPDF_Page_GetGrayFill(page.ptr))
}

func (page *Page) GetGrayStroke() float32 {
	return float32(C.HPDF_Page_GetGrayStroke(page.ptr))
}

const COLOR_SPACE_DEVICE_GRAY int = 0
const COLOR_SPACE_DEVICE_RGB int = 1
const COLOR_SPACE_DEVICE_CMYK int = 2
const COLOR_SPACE_CAL_GRAY int = 3
const COLOR_SPACE_CAL_RGB int = 4
const COLOR_SPACE_LAB int = 5
const COLOR_SPACE_ICC_BASED int = 6
const COLOR_SPACE_SEPARATION int = 7
const COLOR_SPACE_DEVICE_N int = 8
const COLOR_SPACE_INDEXED int = 9
const COLOR_SPACE_PATTERN int = 10

func (page *Page) GetStrokingColorSpace() int {
	return int(C.HPDF_Page_GetStrokingColorSpace(page.ptr))
}

func (page *Page) GetFillingColorSpace() int {
	return int(C.HPDF_Page_GetFillingColorSpace(page.ptr))
}

func (page *Page) GetTextMatrix() TransMatrix {
	rc := C.HPDF_Page_GetTextMatrix(page.ptr)
	return rc.toGo()
}

func (page *Page) GetGStateDepth() uint {
	return uint(C.HPDF_Page_GetGStateDepth(page.ptr))
}

func (page *Page) SetLineWidth(line_width float32) {
	C.HPDF_Page_SetLineWidth(page.ptr, C.float(line_width))
}

func (page *Page) SetLineCap(line_cap int) {
	C.HPDF_Page_SetLineCap(page.ptr, C.HPDF_LineCap(line_cap))
}

func (page *Page) SetLineJoin(line_join int) {
	C.HPDF_Page_SetLineJoin(page.ptr, C.HPDF_LineJoin(line_join))
}

func (page *Page) SetMiterLimit(miter_limit float32) {
	C.HPDF_Page_SetMiterLimit(page.ptr, C.float(miter_limit))
}

func (page *Page) SetDash(mode DashMode) {
	hm := mode.toHaru()
	C.HPDF_Page_SetDash(page.ptr, &hm.ptn[0], hm.num_ptn, hm.phase)
}

func (page *Page) SetFlat(flatness float32) {
	C.HPDF_Page_SetFlat(page.ptr, C.float(flatness))
}

func (page *Page) GSave() {
	C.HPDF_Page_GSave(page.ptr)
}

func (page *Page) GRestore() {
	C.HPDF_Page_GRestore(page.ptr)
}

func (page *Page) Concat(a float32, b float32, c float32, d float32, x float32, y float32) {
	C.HPDF_Page_Concat(page.ptr, C.float(a), C.float(b), C.float(c), C.float(d), C.float(x), C.float(y))
}

func (page *Page) MoveTo(x float32, y float32) {
	C.HPDF_Page_MoveTo(page.ptr, C.float(x), C.float(y))
}

func (page *Page) LineTo(x float32, y float32) {
	C.HPDF_Page_LineTo(page.ptr, C.float(x), C.float(y))
}

func (page *Page) CurveTo(x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32) {
	C.HPDF_Page_CurveTo(page.ptr, C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.float(x3), C.float(y3))
}

func (page *Page) CurveTo2(x2 float32, y2 float32, x3 float32, y3 float32) {
	C.HPDF_Page_CurveTo2(page.ptr, C.float(x2), C.float(y2), C.float(x3), C.float(y3))
}

func (page *Page) CurveTo3(x1 float32, y1 float32, x3 float32, y3 float32) {
	C.HPDF_Page_CurveTo3(page.ptr, C.float(x1), C.float(y1), C.float(x3), C.float(y3))
}

func (page *Page) ClosePath() {
	C.HPDF_Page_ClosePath(page.ptr)
}

func (page *Page) Rectangle(x float32, y float32, width float32, height float32) {
	C.HPDF_Page_Rectangle(page.ptr, C.float(x), C.float(y), C.float(width), C.float(height))
}

func (page *Page) Stroke() {
	C.HPDF_Page_Stroke(page.ptr)
}

func (page *Page) ClosePathStroke() {
	C.HPDF_Page_ClosePathStroke(page.ptr)
}

func (page *Page) Fill() {
	C.HPDF_Page_Fill(page.ptr)
}

func (page *Page) Eofill() {
	C.HPDF_Page_Eofill(page.ptr)
}

func (page *Page) FillStroke() {
	C.HPDF_Page_FillStroke(page.ptr)
}

func (page *Page) EofillStroke() {
	C.HPDF_Page_EofillStroke(page.ptr)
}

func (page *Page) ClosePathFillStroke() {
	C.HPDF_Page_ClosePathFillStroke(page.ptr)
}

func (page *Page) ClosePathEofillStroke() {
	C.HPDF_Page_ClosePathEofillStroke(page.ptr)
}

func (page *Page) EndPath() {
	C.HPDF_Page_EndPath(page.ptr)
}

func (page *Page) Clip() {
	C.HPDF_Page_Clip(page.ptr)
}

func (page *Page) Eoclip() {
	C.HPDF_Page_Eoclip(page.ptr)
}

func (page *Page) BeginText() {
	C.HPDF_Page_BeginText(page.ptr)
}

func (page *Page) EndText() {
	C.HPDF_Page_EndText(page.ptr)
}

func (page *Page) SetCharSpace(value float32) {
	C.HPDF_Page_SetCharSpace(page.ptr, C.float(value))
}

func (page *Page) SetWordSpace(value float32) {
	C.HPDF_Page_SetWordSpace(page.ptr, C.float(value))
}

func (page *Page) SetHorizontalScalling(value float32) {
	C.HPDF_Page_SetHorizontalScalling(page.ptr, C.float(value))
}

func (page *Page) SetTextLeading(value float32) {
	C.HPDF_Page_SetTextLeading(page.ptr, C.float(value))
}

func (page *Page) SetFontAndSize(font Font, size float32) {
	C.HPDF_Page_SetFontAndSize(page.ptr, font.ptr, C.float(size))
}

func (page *Page) SetTextRenderingMode(mode int) {
	C.HPDF_Page_SetTextRenderingMode(page.ptr, C.HPDF_TextRenderingMode(mode))
}

func (page *Page) SetTextRise(value float32) {
	C.HPDF_Page_SetTextRise(page.ptr, C.float(value))
}

func (page *Page) SetTextRaise(value float32) {
	C.HPDF_Page_SetTextRaise(page.ptr, C.float(value))
}

func (page *Page) MoveTextPos(x float32, y float32) {
	C.HPDF_Page_MoveTextPos(page.ptr, C.float(x), C.float(y))
}

func (page *Page) MoveTextPos2(x float32, y float32) {
	C.HPDF_Page_MoveTextPos2(page.ptr, C.float(x), C.float(y))
}

func (page *Page) SetTextMatrix(m TransMatrix) {
	C.HPDF_Page_SetTextMatrix(page.ptr, C.float(m.A), C.float(m.B), C.float(m.C), C.float(m.D), C.float(m.X), C.float(m.Y))

}

func (page *Page) MoveToNextLine() {
	C.HPDF_Page_MoveToNextLine(page.ptr)
}

func (page *Page) ShowText(text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.HPDF_Page_ShowText(page.ptr, _text)
}

func (page *Page) ShowTextNextLine(text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.HPDF_Page_ShowTextNextLine(page.ptr, _text)
}

func (page *Page) ShowTextNextLineEx(word_space float32, char_space float32, text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.HPDF_Page_ShowTextNextLineEx(page.ptr, C.float(word_space), C.float(char_space), _text)

}

func (page *Page) SetGrayFill(gray float32) {
	C.HPDF_Page_SetGrayFill(page.ptr, C.float(gray))
}

func (page *Page) SetGrayStroke(gray float32) {
	C.HPDF_Page_SetGrayStroke(page.ptr, C.float(gray))
}

func (page *Page) SetRGBFill(c RGBColor) {
	C.HPDF_Page_SetRGBFill(page.ptr, C.float(c.R), C.float(c.G), C.float(c.B))
}

func (page *Page) SetRGBStroke(c RGBColor) {
	C.HPDF_Page_SetRGBStroke(page.ptr, C.float(c.R), C.float(c.G), C.float(c.B))
}

func (page *Page) SetCMYKFill(c CMYKColor) {
	C.HPDF_Page_SetCMYKFill(page.ptr, C.float(c.C), C.float(c.M), C.float(c.Y), C.float(c.K))
}

func (page *Page) SetCMYKStroke(c CMYKColor) {
	C.HPDF_Page_SetCMYKStroke(page.ptr, C.float(c.C), C.float(c.M), C.float(c.Y), C.float(c.K))
}

func (page *Page) DrawImage(image Image, x float32, y float32, width float32, height float32) {
	C.HPDF_Page_DrawImage(page.ptr, image.ptr, C.float(x), C.float(y), C.float(width), C.float(height))
}

func (page *Page) Circle(x float32, y float32, ray float32) {
	C.HPDF_Page_Circle(page.ptr, C.float(x), C.float(y), C.float(ray))
}

func (page *Page) Ellipse(x float32, y float32, xray float32, yray float32) {
	C.HPDF_Page_Ellipse(page.ptr, C.float(x), C.float(y), C.float(xray), C.float(yray))
}

func (page *Page) Arc(x float32, y float32, ray float32, ang1 float32, ang2 float32) {
	C.HPDF_Page_Arc(page.ptr, C.float(x), C.float(y), C.float(ray), C.float(ang1), C.float(ang2))
}

func (page *Page) TextOut(xpos float32, ypos float32, text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.HPDF_Page_TextOut(page.ptr, C.float(xpos), C.float(ypos), _text)
}

const TEXT_ALIGN_LEFT int = 0
const TEXT_ALIGN_RIGHT int = 1
const TEXT_ALIGN_CENTER int = 2
const TEXT_ALIGN_JUSTIFY int = 3

func (page *Page) TextRect(left float32, top float32, right float32, bottom float32, text string, align int) uint32 {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	var len C.uint
	C.HPDF_Page_TextRect(page.ptr, C.float(left), C.float(top), C.float(right), C.float(bottom), _text, C.HPDF_TextAlignment(align), &len)
	return uint32(len)
}

const TS_WIPE_RIGHT int = 0
const TS_WIPE_UP int = 1
const TS_WIPE_LEFT int = 2
const TS_WIPE_DOWN int = 3
const TS_BARN_DOORS_HORIZONTAL_OUT int = 4
const TS_BARN_DOORS_HORIZONTAL_IN int = 5
const TS_BARN_DOORS_VERTICAL_OUT int = 6
const TS_BARN_DOORS_VERTICAL_IN int = 7
const TS_BOX_OUT int = 8
const TS_BOX_IN int = 8
const TS_BLINDS_HORIZONTAL int = 10
const TS_BLINDS_VERTICAL int = 11
const TS_DISSOLVE int = 12
const TS_GLITTER_RIGHT int = 13
const TS_GLITTER_DOWN int = 14
const TS_GLITTER_TOP_LEFT_TO_BOTTOM_RIGHT int = 15
const TS_REPLACE int = 16
const TS_EOF int = 17

func (page *Page) SetSlideShow(slideShowType int, disp_time float32, trans_time float32) {
	C.HPDF_Page_SetSlideShow(page.ptr, C.HPDF_TransitionStyle(slideShowType), C.float(disp_time), C.float(trans_time))
}

const PAGE_SIZE_LETTER int = 0
const PAGE_SIZE_LEGAL int = 1
const PAGE_SIZE_A3 int = 2
const PAGE_SIZE_A4 int = 3
const PAGE_SIZE_A5 int = 4
const PAGE_SIZE_B4 int = 5
const PAGE_SIZE_B5 int = 6
const PAGE_SIZE_EXECUTIVE int = 7
const PAGE_SIZE_US4x6 int = 8
const PAGE_SIZE_US4x8 int = 9
const PAGE_SIZE_US5x7 int = 10
const PAGE_SIZE_COMM10 int = 11

const PAGE_PORTRAIT int = 0
const PAGE_LANDSCAPE int = 1

func (page *Page) SetSize(size int, orientation int) {
	C.HPDF_Page_SetSize(page.ptr, C.HPDF_PageSizes(size), C.HPDF_PageDirection(orientation))
}

func (page *Page) TextWidth1(text string, encoding string) (float32, error) {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return 0, e
	}
	defer C.free(unsafe.Pointer(v))
	return float32(C.HPDF_Page_TextWidth(page.ptr, v)), nil
}

func (page *Page) TextOut1(xpos float32, ypos float32, text string, encoding string) error {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return e
	}
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_TextOut(page.ptr, C.float(xpos), C.float(ypos), v)
	return nil
}

func (page *Page) MeasureText1(text string, encoding string, width float32, wordwrap bool) (uint, float32, error) {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return 0, 0, e
	}
	defer C.free(unsafe.Pointer(v))

	var _wordwrap C.int
	if wordwrap {
		_wordwrap = C.HPDF_TRUE
	} else {
		_wordwrap = C.HPDF_FALSE
	}
	var _real_width C.float = 0

	rc := uint(C.HPDF_Page_MeasureText(page.ptr, v, C.float(width), _wordwrap, &_real_width))
	return rc, float32(_real_width), nil
}

func (page *Page) ShowText1(text string, encoding string) error {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return e
	}
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowText(page.ptr, v)
	return nil
}

func (page *Page) ShowTextNextLine1(text string, encoding string) error {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return e
	}
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowTextNextLine(page.ptr, v)
	return nil
}

func (page *Page) ShowTextNextLineEx1(word_space float32, char_space float32, text string, encoding string) error {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return e
	}
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowTextNextLineEx(page.ptr, C.float(word_space), C.float(char_space), v)
	return nil
}

func (page *Page) TextWidth2(text []byte) float32 {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	return float32(C.HPDF_Page_TextWidth(page.ptr, v))
}

func (page *Page) TextOut2(xpos float32, ypos float32, text []byte) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_TextOut(page.ptr, C.float(xpos), C.float(ypos), v)
}

func (page *Page) MeasureText2(text []byte, width float32, wordwrap bool) (uint, float32) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	var _wordwrap C.int
	if wordwrap {
		_wordwrap = C.HPDF_TRUE
	} else {
		_wordwrap = C.HPDF_FALSE
	}
	var _real_width C.float = 0

	rc := uint(C.HPDF_Page_MeasureText(page.ptr, v, C.float(width), _wordwrap, &_real_width))
	return rc, float32(_real_width)
}

func (page *Page) ShowText2(text []byte) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowText(page.ptr, v)
}

func (page *Page) ShowTextNextLine2(text []byte) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowTextNextLine(page.ptr, v)
}

func (page *Page) ShowTextNextLineEx2(word_space float32, char_space float32, text []byte) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowTextNextLineEx(page.ptr, C.float(word_space), C.float(char_space), v)
}
