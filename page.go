package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//Page struct keeps reference to a page of a document
type Page struct {
	ptr C.HPDF_Page
}

//AddPage creates a new page and adds it after the last page of a document.
func (pdf *Doc) AddPage() Page {
	rc := C.HPDF_AddPage(pdf.ptr)
	return Page{ptr: rc}
}

//InsertPage creates a new page and inserts it just before the specified page.
func (pdf *Doc) InsertPage(page Page) Page {
	rc := C.HPDF_InsertPage(pdf.ptr, page.ptr)
	return Page{ptr: rc}
}

//GetCurrentPage gets current page of the document
func (pdf *Doc) GetCurrentPage() Page {
	rc := C.HPDF_GetCurrentPage(pdf.ptr)
	return Page{ptr: rc}
}

//GetPageByIndex gets page of the document by its index
func (pdf *Doc) GetPageByIndex(index uint) Page {
	rc := C.HPDF_GetPageByIndex(pdf.ptr, C.uint(index))
	return Page{ptr: rc}
}

//SetWidth changes the width of a page.
func (page *Page) SetWidth(value float32) {
	C.HPDF_Page_SetWidth(page.ptr, C.float(value))
}

//SetHeight changes the height of a page.
func (page *Page) SetHeight(value float32) {
	C.HPDF_Page_SetHeight(page.ptr, C.float(value))
}

//SetRotate sets rotation angle of the page.
func (page *Page) SetRotate(angle uint16) {
	C.HPDF_Page_SetRotate(page.ptr, C.ushort(angle))
}

//SetZoom sets default page zoom
func (page *Page) SetZoom(zoom float32) {
	C.HPDF_Page_SetZoom(page.ptr, C.float(zoom))
}

//PageNumStyleDecimal is page numeration style
const PageNumStyleDecimal int = 0 
//PageNumStyleUpperRoman is page numeration style
const PageNumStyleUpperRoman int = 1 
//PageNumStyleLowerRoman is page numeration style
const PageNumStyleLowerRoman int = 2 
//PageNumStyleUpperLetters is page numeration style
const PageNumStyleUpperLetters int = 3 
//PageNumStyleLowerLetters is page numeration style
const PageNumStyleLowerLetters int = 4 

//AddPageLabel adds a page labeling range for the document.
//
//Use PageNumStyle* values as a value for style
func (pdf *Doc) AddPageLabel(pageNum uint, style int, firstPage uint, prefix string) {
	_prefix := C.CString(prefix)
	defer C.free(unsafe.Pointer(_prefix))
	C.HPDF_AddPageLabel(pdf.ptr, C.uint(pageNum), C.HPDF_PageNumStyle(style), C.uint(firstPage), _prefix)
}

//PageLayoutSingle is page layout style
const PageLayoutSingle int = 0 
//PageLayoutOneColumn is page layout style
const PageLayoutOneColumn int = 1 
//PageLayoutTwoColumnLeft is page layout style
const PageLayoutTwoColumnLeft int = 2 
//PageLayoutTwoColumnRight is page layout style
const PageLayoutTwoColumnRight int = 3 
//PageLayoutTwoPageLeft is page layout style
const PageLayoutTwoPageLeft int = 4 
//PageLayoutTwoPageRight is page layout style
const PageLayoutTwoPageRight int = 5 

//GetPageLayout returns the handle of current page object.
//
//The function returns PageLayout* value
func (pdf *Doc) GetPageLayout() int {
	rc := C.HPDF_GetPageLayout(pdf.ptr)
	return int(rc)
}

//SetPageLayout sets how the page should be displayed. If this attribute is not set, the setting of the viewer application is used.
//
//Use PageLayout as layout value
func (pdf *Doc) SetPageLayout(layout int) {
	C.HPDF_SetPageLayout(pdf.ptr, C.HPDF_PageLayout(layout))
}

//PageModeUseNone is page use mode
const PageModeUseNone int = 0 
//PageModeUseOutline is page use mode
const PageModeUseOutline int = 1 
//PageModeUseThumbs is page use mode
const PageModeUseThumbs int = 2 
//PageModeFullScreen is page use mode
const PageModeFullScreen int = 3 

//GetPageMode returns page mode
//
//The function return PageNode* value
func (pdf *Doc) GetPageMode() int {
	rc := C.HPDF_GetPageMode(pdf.ptr)
	return int(rc)
}

//SetPageMode sets how the document should be displayed.
func (pdf *Doc) SetPageMode(mode int) {
	C.HPDF_SetPageMode(pdf.ptr, C.HPDF_PageMode(mode))
}

//TextWidth gets the width of the text in current fontsize, character spacing and word spacing.
func (page *Page) TextWidth(text string) float32 {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	return float32(C.HPDF_Page_TextWidth(page.ptr, _text))
}

//MeasureText calculates the byte length which can be included within the specified width.
func (page *Page) MeasureText(text string, width float32, wordwrap bool) (uint, float32) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	var _wordwrap C.int
	if wordwrap {
		_wordwrap = C.HPDF_TRUE
	} else {
		_wordwrap = C.HPDF_FALSE
	}
	var _realWidth C.float

	rc := uint(C.HPDF_Page_MeasureText(page.ptr, _text, C.float(width), _wordwrap, &_realWidth))
	return rc, float32(_realWidth)
}

//GetWidth gets the width of the page.
func (page *Page) GetWidth() float32 {
	return float32(C.HPDF_Page_GetWidth(page.ptr))
}

//GetHeight gets the height of a page.
func (page *Page) GetHeight() float32 {
	return float32(C.HPDF_Page_GetHeight(page.ptr))
}

//GetGMode gets the current graphics mode.
func (page *Page) GetGMode() uint16 {
	return uint16(C.HPDF_Page_GetGMode(page.ptr))
}

//GetCurrentPos gets the current position for path painting.
func (page *Page) GetCurrentPos() Point {
	rc := C.HPDF_Page_GetCurrentPos(page.ptr)
	return Point{X: float32(rc.x), Y: float32(rc.y)}

}

//GetCurrentTextPos gets the current position for text showing.
func (page *Page) GetCurrentTextPos() Point {
	rc := C.HPDF_Page_GetCurrentTextPos(page.ptr)
	return Point{X: float32(rc.x), Y: float32(rc.y)}
}

//GetCurrentFont gets the handle of the page's current font.
func (page *Page) GetCurrentFont() Font {
	rc := C.HPDF_Page_GetCurrentFont(page.ptr)
	return Font{ptr: rc}
}

//GetCurrentFontSize gets the size of the page's current font.
func (page *Page) GetCurrentFontSize() float32 {
	return float32(C.HPDF_Page_GetCurrentFontSize(page.ptr))
}

//GetTransMatrix gets the current transformation matrix of the page.
func (page *Page) GetTransMatrix() TransMatrix {
	rc := C.HPDF_Page_GetTransMatrix(page.ptr)
	return rc.toGo()
}

//GetLineWidth gets the current line width of the page.
func (page *Page) GetLineWidth() float32 {
	return float32(C.HPDF_Page_GetLineWidth(page.ptr))
}

//ButtEnd is line end style
const ButtEnd int = 0 
//RoundEnd is line end style
const RoundEnd int = 1 
//ProjectingSquareEnd is line end style
const ProjectingSquareEnd int = 2 
//MiterJoin is line join style
const MiterJoin int = 0 
//RoundJoin is line join style
const RoundJoin int = 1 
//BevelJoin is line join style
const BevelJoin int = 2 

//GetLineCap gets the current line cap style of the page.
//
//the method return Buttend, RoundEnd, ProjectingSquareEnd
func (page *Page) GetLineCap() int {
	return int(C.HPDF_Page_GetLineCap(page.ptr))
}

//GetLineJoin gets the current line join style of the page.
//
//the method return MiterJoin, RoundJoin or BevelJoin
func (page *Page) GetLineJoin() int {
	return int(C.HPDF_Page_GetLineJoin(page.ptr))
}

//GetMiterLimit gets the current value of the page's miter limit.
func (page *Page) GetMiterLimit() float32 {
	return float32(C.HPDF_Page_GetMiterLimit(page.ptr))
}

//GetDash gets the current pattern of the page.
func (page *Page) GetDash() DashMode {
	rc := C.HPDF_Page_GetDash(page.ptr)
	return rc.toGo()
}

//GetFlat gets the current value of the page's flatness.
func (page *Page) GetFlat() float32 {
	return float32(C.HPDF_Page_GetFlat(page.ptr))
}

//GetCharSpace gets the current value of the page's character spacing.
func (page *Page) GetCharSpace() float32 {
	return float32(C.HPDF_Page_GetCharSpace(page.ptr))
}

//GetWordSpace returns the current value of the page's word spacing.
func (page *Page) GetWordSpace() float32 {
	return float32(C.HPDF_Page_GetWordSpace(page.ptr))
}

//GetHorizontalScalling returns the current value of the page's horizontal scalling for text showing.
func (page *Page) GetHorizontalScalling() float32 {
	return float32(C.HPDF_Page_GetHorizontalScalling(page.ptr))
}

//GetTextLeading returns the current value of the page's line spacing.
func (page *Page) GetTextLeading() float32 {
	return float32(C.HPDF_Page_GetTextLeading(page.ptr))
}

//TextRenderFill is text rendering mode
const TextRenderFill int = 0 
//TextRenderStroke is text rendering mode
const TextRenderStroke int = 1 
//TextRenderFillThenStroke is text rendering mode
const TextRenderFillThenStroke int = 2 
//TextRenderInvisible is text rendering mode
const TextRenderInvisible int = 3 
//TextRenderFillClipping is text rendering mode
const TextRenderFillClipping int = 4 
//TextRenderStrokeClipping is text rendering mode
const TextRenderStrokeClipping int = 5 
//TextRenderFillStrokeClipping is text rendering mode
const TextRenderFillStrokeClipping int = 6 
//TextRenderClipping is text rendering mode
const TextRenderClipping int = 7 

//GetTextRenderingMode returns the current value of the page's text rendering mode.
//
//The method returns TextRender* value
func (page *Page) GetTextRenderingMode() int {
	return int(C.HPDF_Page_GetTextRenderingMode(page.ptr))
}

//GetTextRaise returns text raise
func (page *Page) GetTextRaise() float32 {
	return float32(C.HPDF_Page_GetTextRaise(page.ptr))
}

//GetTextRise returns the current value of the page's text rising.
func (page *Page) GetTextRise() float32 {
	return float32(C.HPDF_Page_GetTextRise(page.ptr))
}

//GetRGBFill returns the current value of the page's filling color. HPDF_Page_GetRGBFill() is valid only when the page's filling color space is HPDF_CS_DEVICE_RGB.
func (page *Page) GetRGBFill() RGBColor {
	rc := C.HPDF_Page_GetRGBFill(page.ptr)
	return rc.toGo()
}

//GetRGBStroke returns the current value of the page's stroking color. HPDF_Page_GetRGBStroke() is valid only when the page's stroking color space is HPDF_CS_DEVICE_RGB.
func (page *Page) GetRGBStroke() RGBColor {
	rc := C.HPDF_Page_GetRGBStroke(page.ptr)
	return rc.toGo()
}

//GetCMYKFill returns the current value of the page's filling color. HPDF_Page_GetCMYKFill() is valid only when the page's filling color space is HPDF_CS_DEVICE_CMYK.
func (page *Page) GetCMYKFill() CMYKColor {
	rc := C.HPDF_Page_GetCMYKFill(page.ptr)
	return rc.toGo()
}

//GetCMYKStroke returns the current value of the page's stroking color. HPDF_Get_CMYKStroke() is valid only when the page's stroking color space is HPDF_CS_DEVICE_CMYK.
func (page *Page) GetCMYKStroke() CMYKColor {
	rc := C.HPDF_Page_GetCMYKStroke(page.ptr)
	return rc.toGo()

}

//GetGrayFill returns the current value of the page's filling color. HPDF_Page_GetGrayFill() is valid only when the page's stroking color space is HPDF_CS_DEVICE_GRAY.
func (page *Page) GetGrayFill() float32 {
	return float32(C.HPDF_Page_GetGrayFill(page.ptr))
}

//GetGrayStroke returns the current value of the page's stroking color. HPDF_Page_GetGrayStroke() is valid only when the page's stroking color space is HPDF_CS_DEVICE_GRAY.
func (page *Page) GetGrayStroke() float32 {
	return float32(C.HPDF_Page_GetGrayStroke(page.ptr))
}

//ColorSpaceDeviceGray is color space identifier
const ColorSpaceDeviceGray int = 0 
//ColorSpaceDeviceRgb is color space identifier
const ColorSpaceDeviceRgb int = 1 
//ColorSpaceDeviceCmyk is color space identifier
const ColorSpaceDeviceCmyk int = 2 
//ColorSpaceCalGray is color space identifier
const ColorSpaceCalGray int = 3 
//ColorSpaceCalRgb is color space identifier
const ColorSpaceCalRgb int = 4 
//ColorSpaceLab is color space identifier
const ColorSpaceLab int = 5 
//ColorSpaceIccBased is color space identifier
const ColorSpaceIccBased int = 6 
//ColorSpaceSeparation is color space identifier
const ColorSpaceSeparation int = 7 
//ColorSpaceDeviceN is color space identifier
const ColorSpaceDeviceN int = 8 
//ColorSpaceIndexed is color space identifier
const ColorSpaceIndexed int = 9 
//ColorSpacePattern is color space identifier
const ColorSpacePattern int = 10 

//GetStrokingColorSpace returns the current value of the page's stroking color space.
//
//The methgod return ColorSpace* value
func (page *Page) GetStrokingColorSpace() int {
	return int(C.HPDF_Page_GetStrokingColorSpace(page.ptr))
}

//GetFillingColorSpace returns the current value of the page's stroking color space.
//
//The methgod return ColorSpace* value
func (page *Page) GetFillingColorSpace() int {
	return int(C.HPDF_Page_GetFillingColorSpace(page.ptr))
}

//GetTextMatrix gets the current text transformation matrix of the page.
func (page *Page) GetTextMatrix() TransMatrix {
	rc := C.HPDF_Page_GetTextMatrix(page.ptr)
	return rc.toGo()
}

//GetGStateDepth returns the number of the page's graphics state stack.
func (page *Page) GetGStateDepth() uint {
	return uint(C.HPDF_Page_GetGStateDepth(page.ptr))
}

//SetLineWidth sets the width of the line used to stroke a path.
func (page *Page) SetLineWidth(lineWidth float32) {
	C.HPDF_Page_SetLineWidth(page.ptr, C.float(lineWidth))
}

//SetLineCap sets the shape to be used at the ends of lines.
//
//Use ButtEnd, RoundEnd, ProjectingSquareEnd as cap style
func (page *Page) SetLineCap(lineCap int) {
	C.HPDF_Page_SetLineCap(page.ptr, C.HPDF_LineCap(lineCap))
}

//SetLineJoin Sets the line join style in the page.
//
//Use MiterJoin, RoundJoin or BevelJoin as join style
func (page *Page) SetLineJoin(lineJoin int) {
	C.HPDF_Page_SetLineJoin(page.ptr, C.HPDF_LineJoin(lineJoin))
}

//SetMiterLimit sets miter limit
func (page *Page) SetMiterLimit(miterLimit float32) {
	C.HPDF_Page_SetMiterLimit(page.ptr, C.float(miterLimit))
}

//SetDash sets the dash pattern for lines in the page.
func (page *Page) SetDash(mode DashMode) {
	hm := mode.toHaru()
	C.HPDF_Page_SetDash(page.ptr, &hm.ptn[0], hm.num_ptn, hm.phase)
}

//SetFlat sets flatness
func (page *Page) SetFlat(flatness float32) {
	C.HPDF_Page_SetFlat(page.ptr, C.float(flatness))
}

//GSave saves the page's current graphics parameter to the stack. An application can invoke HPDF_Page_GSave() up to 28 (???) and can restore the saved parameter by invoking HPDF_Page_GRestore().
func (page *Page) GSave() {
	C.HPDF_Page_GSave(page.ptr)
}

//GRestore restore the graphics state which is saved by HPDF_Page_GSave().
func (page *Page) GRestore() {
	C.HPDF_Page_GRestore(page.ptr)
}

//Concat concatenates the page's current transformation matrix and specified matrix.
func (page *Page) Concat(a float32, b float32, c float32, d float32, x float32, y float32) {
	C.HPDF_Page_Concat(page.ptr, C.float(a), C.float(b), C.float(c), C.float(d), C.float(x), C.float(y))
}

//MoveTo starts a new subpath and move the current point for drawing path,  HPDF_Page_MoveTo() sets the start point for the path to the point (x, y).
func (page *Page) MoveTo(x float32, y float32) {
	C.HPDF_Page_MoveTo(page.ptr, C.float(x), C.float(y))
}

//LineTo appends a path from the current point to the specified point.
func (page *Page) LineTo(x float32, y float32) {
	C.HPDF_Page_LineTo(page.ptr, C.float(x), C.float(y))
}

//CurveTo appends a Bezier curve to the current path using the control points (x1, y1) and (x2, y2) and (x3, y3), then sets the current point to (x3, y3).
func (page *Page) CurveTo(x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32) {
	C.HPDF_Page_CurveTo(page.ptr, C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.float(x3), C.float(y3))
}

//CurveTo2 appends a Bezier curve to the current path using the current point and (x2, y2) and (x3, y3) as control points. Then, the current point is set to (x3, y3).
func (page *Page) CurveTo2(x2 float32, y2 float32, x3 float32, y3 float32) {
	C.HPDF_Page_CurveTo2(page.ptr, C.float(x2), C.float(y2), C.float(x3), C.float(y3))
}

//CurveTo3 appends a Bezier curve to the current path using two specified points.
func (page *Page) CurveTo3(x1 float32, y1 float32, x3 float32, y3 float32) {
	C.HPDF_Page_CurveTo3(page.ptr, C.float(x1), C.float(y1), C.float(x3), C.float(y3))
}

//ClosePath appends a straight line from the current point to the start point of sub path.
func (page *Page) ClosePath() {
	C.HPDF_Page_ClosePath(page.ptr)
}

//Rectangle appends a rectangle to the current path.
func (page *Page) Rectangle(x float32, y float32, width float32, height float32) {
	C.HPDF_Page_Rectangle(page.ptr, C.float(x), C.float(y), C.float(width), C.float(height))
}

//Stroke paints the current path.
func (page *Page) Stroke() {
	C.HPDF_Page_Stroke(page.ptr)
}

//ClosePathStroke closes the current path.
func (page *Page) ClosePathStroke() {
	C.HPDF_Page_ClosePathStroke(page.ptr)
}

//Fill fills the current path using the nonzero winding number rule.
func (page *Page) Fill() {
	C.HPDF_Page_Fill(page.ptr)
}

//Eofill fills the current path using the even-odd rule.
func (page *Page) Eofill() {
	C.HPDF_Page_Eofill(page.ptr)
}

//FillStroke fills the current path using the nonzero winding number rule, then paints the path.
func (page *Page) FillStroke() {
	C.HPDF_Page_FillStroke(page.ptr)
}

//EofillStroke fills the current path using the even-odd rule, then paints the path.
func (page *Page) EofillStroke() {
	C.HPDF_Page_EofillStroke(page.ptr)
}

//ClosePathFillStroke closes the current path, fills the current path using the nonzero winding number rule, then paints the path.
func (page *Page) ClosePathFillStroke() {
	C.HPDF_Page_ClosePathFillStroke(page.ptr)
}

//ClosePathEofillStroke closes the current path, fills the current path using the even-odd rule, then paints the path.
func (page *Page) ClosePathEofillStroke() {
	C.HPDF_Page_ClosePathEofillStroke(page.ptr)
}

//EndPath ends the path object without filling or painting.
func (page *Page) EndPath() {
	C.HPDF_Page_EndPath(page.ptr)
}

//Clip modifies the current clipping path by intersecting it with the current path using the even-odd rule. The clipping path is only modified after the succeeding painting operator. To avoid painting the current path, use the function HPDF_Page_EndPath().
func (page *Page) Clip() {
	C.HPDF_Page_Clip(page.ptr)
}

//Eoclip  ???
func (page *Page) Eoclip() {
	C.HPDF_Page_Eoclip(page.ptr)
}

//BeginText begins a text object and sets the text position to (0, 0).
func (page *Page) BeginText() {
	C.HPDF_Page_BeginText(page.ptr)
}

//EndText ends a text object.
func (page *Page) EndText() {
	C.HPDF_Page_EndText(page.ptr)
}

//SetCharSpace sets the character spacing for text.
func (page *Page) SetCharSpace(value float32) {
	C.HPDF_Page_SetCharSpace(page.ptr, C.float(value))
}

//SetWordSpace sets the word spacing for text.
func (page *Page) SetWordSpace(value float32) {
	C.HPDF_Page_SetWordSpace(page.ptr, C.float(value))
}

//SetHorizontalScalling sets the horizontal scalling (scaling) for text showing.
func (page *Page) SetHorizontalScalling(value float32) {
	C.HPDF_Page_SetHorizontalScalling(page.ptr, C.float(value))
}

//SetTextLeading sets the text leading (line spacing) for text showing.
func (page *Page) SetTextLeading(value float32) {
	C.HPDF_Page_SetTextLeading(page.ptr, C.float(value))
}

//SetFontAndSize sets the type of font and size leading.
func (page *Page) SetFontAndSize(font Font, size float32) {
	C.HPDF_Page_SetFontAndSize(page.ptr, font.ptr, C.float(size))
}

//SetTextRenderingMode sets the text rendering mode.
//
//Use TextRender* as a mode
func (page *Page) SetTextRenderingMode(mode int) {
	C.HPDF_Page_SetTextRenderingMode(page.ptr, C.HPDF_TextRenderingMode(mode))
}

//SetTextRise moves the text position in vertical direction by
func (page *Page) SetTextRise(value float32) {
	C.HPDF_Page_SetTextRise(page.ptr, C.float(value))
}

//SetTextRaise sets the text raise
func (page *Page) SetTextRaise(value float32) {
	C.HPDF_Page_SetTextRaise(page.ptr, C.float(value))
}

//MoveTextPos changes the current text position, using the specified offset values. If the current text position is (x1, y1), the new text position will be (x1 + x, y1 + y).
func (page *Page) MoveTextPos(x float32, y float32) {
	C.HPDF_Page_MoveTextPos(page.ptr, C.float(x), C.float(y))
}

//MoveTextPos2 changes the current text position, using the specified offset values. If the current text position is (x1, y1), the new text position will be (x1 + x, y1 + y).
func (page *Page) MoveTextPos2(x float32, y float32) {
	C.HPDF_Page_MoveTextPos2(page.ptr, C.float(x), C.float(y))
}

//SetTextMatrix sets text transformation matrix
func (page *Page) SetTextMatrix(m TransMatrix) {
	C.HPDF_Page_SetTextMatrix(page.ptr, C.float(m.A), C.float(m.B), C.float(m.C), C.float(m.D), C.float(m.X), C.float(m.Y))

}

//MoveToNextLine moves current position for the text showing depending on current text showing point and text leading. The new position is calculated with current text transition matrix.
func (page *Page) MoveToNextLine() {
	C.HPDF_Page_MoveToNextLine(page.ptr)
}

//ShowText prints the text at the current position on the page.
func (page *Page) ShowText(text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.HPDF_Page_ShowText(page.ptr, _text)
}

//ShowTextNextLine moves the current text position to the start of the next line, then prints the text at the current position on the page.
func (page *Page) ShowTextNextLine(text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.HPDF_Page_ShowTextNextLine(page.ptr, _text)
}

//ShowTextNextLineEx moves the current text position to the start of the next line; then sets word spacing and character spacing; finally prints the text at the current position on the page.
func (page *Page) ShowTextNextLineEx(wordSpace float32, charSpace float32, text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.HPDF_Page_ShowTextNextLineEx(page.ptr, C.float(wordSpace), C.float(charSpace), _text)

}

//SetGrayFill sets the filling color.
func (page *Page) SetGrayFill(gray float32) {
	C.HPDF_Page_SetGrayFill(page.ptr, C.float(gray))
}

//SetGrayStroke sets the stroking color.
func (page *Page) SetGrayStroke(gray float32) {
	C.HPDF_Page_SetGrayStroke(page.ptr, C.float(gray))
}

//SetRGBFill sets the filling color.
func (page *Page) SetRGBFill(c RGBColor) {
	C.HPDF_Page_SetRGBFill(page.ptr, C.float(c.R), C.float(c.G), C.float(c.B))
}

//SetRGBStroke sets the stroking color.
func (page *Page) SetRGBStroke(c RGBColor) {
	C.HPDF_Page_SetRGBStroke(page.ptr, C.float(c.R), C.float(c.G), C.float(c.B))
}

//SetCMYKFill sets the filling color.
func (page *Page) SetCMYKFill(c CMYKColor) {
	C.HPDF_Page_SetCMYKFill(page.ptr, C.float(c.C), C.float(c.M), C.float(c.Y), C.float(c.K))
}

//SetCMYKStroke sets the stroking color.
func (page *Page) SetCMYKStroke(c CMYKColor) {
	C.HPDF_Page_SetCMYKStroke(page.ptr, C.float(c.C), C.float(c.M), C.float(c.Y), C.float(c.K))
}

//DrawImage shows an image in one operation.
func (page *Page) DrawImage(image Image, x float32, y float32, width float32, height float32) {
	C.HPDF_Page_DrawImage(page.ptr, image.ptr, C.float(x), C.float(y), C.float(width), C.float(height))
}

//Circle appends a circle to the current path.
func (page *Page) Circle(x float32, y float32, ray float32) {
	C.HPDF_Page_Circle(page.ptr, C.float(x), C.float(y), C.float(ray))
}

//Ellipse appends an ellipse to the current path.
func (page *Page) Ellipse(x float32, y float32, xray float32, yray float32) {
	C.HPDF_Page_Ellipse(page.ptr, C.float(x), C.float(y), C.float(xray), C.float(yray))
}

//Arc appends a circle arc to the current path. Angles are given in degrees, with 0 degrees being vertical, upward, from the (x,y) position.
func (page *Page) Arc(x float32, y float32, ray float32, ang1 float32, ang2 float32) {
	C.HPDF_Page_Arc(page.ptr, C.float(x), C.float(y), C.float(ray), C.float(ang1), C.float(ang2))
}

//TextOut prints the text on the specified position.
func (page *Page) TextOut(xpos float32, ypos float32, text string) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	C.HPDF_Page_TextOut(page.ptr, C.float(xpos), C.float(ypos), _text)
}

//TextAlignLeft is text alignment 
const TextAlignLeft int = 0 
//TextAlignRight is text alignment 
const TextAlignRight int = 1 
//TextAlignCenter is text alignment 
const TextAlignCenter int = 2 
//TextAlignJustify is text alignment 
const TextAlignJustify int = 3 

//TextRect draw text in the specified rectangle
//
//Use TextAlign* as a alignment
func (page *Page) TextRect(left float32, top float32, right float32, bottom float32, text string, align int) uint32 {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	var len C.uint
	C.HPDF_Page_TextRect(page.ptr, C.float(left), C.float(top), C.float(right), C.float(bottom), _text, C.HPDF_TextAlignment(align), &len)
	return uint32(len)
}

//TsWipeRight is slide show mode 
const TsWipeRight int = 0 
//TsWipeUp is slide show mode 
const TsWipeUp int = 1 
//TsWipeLeft is slide show mode 
const TsWipeLeft int = 2 
//TsWipeDown is slide show mode 
const TsWipeDown int = 3 
//TsBarnDoorsHorizontalOut is slide show mode 
const TsBarnDoorsHorizontalOut int = 4 
//TsBarnDoorsHorizontalIn is slide show mode 
const TsBarnDoorsHorizontalIn int = 5 
//TsBarnDoorsVerticalOut is slide show mode 
const TsBarnDoorsVerticalOut int = 6 
//TsBarnDoorsVerticalIn is slide show mode 
const TsBarnDoorsVerticalIn int = 7 
//TsBoxOut is slide show mode 
const TsBoxOut int = 8 
//TsBoxIn is slide show mode 
const TsBoxIn int = 8 
//TsBlindsHorizontal is slide show mode 
const TsBlindsHorizontal int = 10 
//TsBlindsVertical is slide show mode 
const TsBlindsVertical int = 11 
//TsDissolve is slide show mode 
const TsDissolve int = 12 
//TsGlitterRight is slide show mode 
const TsGlitterRight int = 13 
//TsGlitterDown is slide show mode 
const TsGlitterDown int = 14 
//TsGlitterTopLeftToBottomRight is slide show mode 
const TsGlitterTopLeftToBottomRight int = 15 
//TsReplace is slide show mode 
const TsReplace int = 16 

//SetSlideShow configures the setting for slide transition of the page.
//
//use Ts* as slideShowType
func (page *Page) SetSlideShow(slideShowType int, dispTime float32, transTime float32) {
	C.HPDF_Page_SetSlideShow(page.ptr, C.HPDF_TransitionStyle(slideShowType), C.float(dispTime), C.float(transTime))
}

//PageSizeLetter is page size 
const PageSizeLetter int = 0 
//PageSizeLegal is page size 
const PageSizeLegal int = 1 
//PageSizeA3 is page size 
const PageSizeA3 int = 2 
//PageSizeA4 is page size 
const PageSizeA4 int = 3 
//PageSizeA5 is page size 
const PageSizeA5 int = 4 
//PageSizeB4 is page size 
const PageSizeB4 int = 5 
//PageSizeB5 is page size 
const PageSizeB5 int = 6 
//PageSizeExecutive is page size 
const PageSizeExecutive int = 7 
//PageSizeUs4x6 is page size 
const PageSizeUs4x6 int = 8 
//PageSizeUs4x8 is page size 
const PageSizeUs4x8 int = 9 
//PageSizeUs5x7 is page size 
const PageSizeUs5x7 int = 10 
//PageSizeComm10 is page size 
const PageSizeComm10 int = 11 

//PagePortrait is page orientation
const PagePortrait int = 0 
//PageLandscape is orientation
const PageLandscape int = 1 

//SetSize changes the size and direction of a page to a predefined size.
//
//Use PageSize* as page size. Use PagePortrait or PageLandscape as orientation
func (page *Page) SetSize(size int, orientation int) {
	C.HPDF_Page_SetSize(page.ptr, C.HPDF_PageSizes(size), C.HPDF_PageDirection(orientation))
}

//TextWidth1 measures the width of the text
func (page *Page) TextWidth1(text string, encoding string) (float32, error) {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return 0, e
	}
	defer C.free(unsafe.Pointer(v))
	return float32(C.HPDF_Page_TextWidth(page.ptr, v)), nil
}

//TextOut1 writes the text
func (page *Page) TextOut1(xpos float32, ypos float32, text string, encoding string) error {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return e
	}
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_TextOut(page.ptr, C.float(xpos), C.float(ypos), v)
	return nil
}

//MeasureText1 measures the text dimensions
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
	var _realWidth C.float

	rc := uint(C.HPDF_Page_MeasureText(page.ptr, v, C.float(width), _wordwrap, &_realWidth))
	return rc, float32(_realWidth), nil
}

//ShowText1 shows text at the current position
func (page *Page) ShowText1(text string, encoding string) error {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return e
	}
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowText(page.ptr, v)
	return nil
}

//ShowTextNextLine1 shows text at the next line
func (page *Page) ShowTextNextLine1(text string, encoding string) error {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return e
	}
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowTextNextLine(page.ptr, v)
	return nil
}

//ShowTextNextLineEx1 shows text at the next line
func (page *Page) ShowTextNextLineEx1(wordSpace float32, charSpace float32, text string, encoding string) error {
	v, _, e := CEncodedString(text, encoding)
	if e != nil {
		return e
	}
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowTextNextLineEx(page.ptr, C.float(wordSpace), C.float(charSpace), v)
	return nil
}

//TextWidth2 measures the text width
func (page *Page) TextWidth2(text []byte) float32 {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	return float32(C.HPDF_Page_TextWidth(page.ptr, v))
}

//TextOut2 writes the text at the specified position
func (page *Page) TextOut2(xpos float32, ypos float32, text []byte) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_TextOut(page.ptr, C.float(xpos), C.float(ypos), v)
}

//MeasureText2 measures the text size
func (page *Page) MeasureText2(text []byte, width float32, wordwrap bool) (uint, float32) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	var _wordwrap C.int
	if wordwrap {
		_wordwrap = C.HPDF_TRUE
	} else {
		_wordwrap = C.HPDF_FALSE
	}
	var _realWidth C.float 

	rc := uint(C.HPDF_Page_MeasureText(page.ptr, v, C.float(width), _wordwrap, &_realWidth))
	return rc, float32(_realWidth)
}

//ShowText2 shows the text at the specified position
func (page *Page) ShowText2(text []byte) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowText(page.ptr, v)
}

//ShowTextNextLine2 shows the text at the next line
func (page *Page) ShowTextNextLine2(text []byte) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowTextNextLine(page.ptr, v)
}

//ShowTextNextLineEx2 shows the text at the next line
func (page *Page) ShowTextNextLineEx2(wordSpace float32, charSpace float32, text []byte) {
	v := BufferToASCIZ(text)
	defer C.free(unsafe.Pointer(v))
	C.HPDF_Page_ShowTextNextLineEx(page.ptr, C.float(wordSpace), C.float(charSpace), v)
}
