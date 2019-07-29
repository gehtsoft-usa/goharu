package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import (
	"time"
	"unsafe"
)

//The structure represents an image
type Annotation struct {
	ptr C.HPDF_Annotation
}

func (page *Page) Page_Create3DAnnot(rect Rect, tb bool, np bool, u3d U3D, ap Image) Annotation {
	var _tb C.int
	if tb {
		_tb = C.HPDF_TRUE
	} else {
		_tb = C.HPDF_FALSE
	}

	var _np C.int
	if np {
		_np = C.HPDF_TRUE
	} else {
		_np = C.HPDF_FALSE
	}

	rc := C.HPDF_Page_Create3DAnnot(page.ptr, rect.toHaru(), _tb, _np, u3d.ptr, ap.ptr)
	return Annotation{ptr: rc}
}

func (page *Page) Page_CreateTextAnnot(rect Rect, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateTextAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)
	return Annotation{ptr: rc}
}

func (page *Page) Page_CreateFreeTextAnnot(rect Rect, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateFreeTextAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)
	return Annotation{ptr: rc}
}

func (page *Page) Page_CreateLineAnnot(text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateLineAnnot(page.ptr, _text, encoder.ptr)
	return Annotation{ptr: rc}
}

func (pdf *Doc) Page_CreateWidgetAnnot_WhiteOnlyWhilePrint(page Page, rect Rect) Annotation {

	rc := C.HPDF_Page_CreateWidgetAnnot_WhiteOnlyWhilePrint(pdf.ptr, page.ptr, rect.toHaru())
	return Annotation{ptr: rc}
}

func (page *Page) Page_CreateWidgetAnnot(rect Rect) Annotation {

	rc := C.HPDF_Page_CreateWidgetAnnot(page.ptr, rect.toHaru())

	return Annotation{ptr: rc}

}

func (page *Page) Page_CreateLinkAnnot(rect Rect, dst Destination) Annotation {
	rc := C.HPDF_Page_CreateLinkAnnot(page.ptr, rect.toHaru(), dst.ptr)
	return Annotation{ptr: rc}

}

func (page *Page) Page_CreateURILinkAnnot(rect Rect, uri string) Annotation {
	_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(_uri))
	rc := C.HPDF_Page_CreateURILinkAnnot(page.ptr, rect.toHaru(), _uri)
	return Annotation{ptr: rc}

}

const ANNOT_TEXT_NOTES int = 0
const ANNOT_LINK int = 1
const ANNOT_SOUND int = 2
const ANNOT_FREE_TEXT int = 3
const ANNOT_STAMP int = 4
const ANNOT_SQUARE int = 5
const ANNOT_CIRCLE int = 6
const ANNOT_STRIKE_OUT int = 7
const ANNOT_HIGHTLIGHT int = 8
const ANNOT_UNDERLINE int = 9
const ANNOT_INK int = 10
const ANNOT_FILE_ATTACHMENT int = 11
const ANNOT_POPUP int = 12
const ANNOT_3D int = 13
const ANNOT_SQUIGGLY int = 14
const ANNOT_LINE int = 15
const ANNOT_PROJECTION int = 16
const ANNOT_WIDGET int = 17

func (page *Page) Page_CreateTextMarkupAnnot(rect Rect, text string, encoder Encoder, subType int) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateTextMarkupAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr, C.HPDF_AnnotType(subType))
	return Annotation{ptr: rc}

}

func (page *Page) Page_CreateHighlightAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateHighlightAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

func (page *Page) Page_CreateUnderlineAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateUnderlineAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

func (page *Page) Page_CreateSquigglyAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateSquigglyAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

func (page *Page) Page_CreateStrikeOutAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateStrikeOutAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

func (page *Page) Page_CreatePopupAnnot(rect Rect, parent Annotation) Annotation {

	rc := C.HPDF_Page_CreatePopupAnnot(page.ptr, rect.toHaru(), parent.ptr)

	return Annotation{ptr: rc}

}

const STAMP_ANNOT_APPROVED int = 0
const STAMP_ANNOT_EXPERIMENTAL int = 1
const STAMP_ANNOT_NOTAPPROVED int = 2
const STAMP_ANNOT_ASIS int = 3
const STAMP_ANNOT_EXPIRED int = 4
const STAMP_ANNOT_NOTFORPUBLICRELEASE int = 5
const STAMP_ANNOT_CONFIDENTIAL int = 6
const STAMP_ANNOT_FINAL int = 7
const STAMP_ANNOT_SOLD int = 8
const STAMP_ANNOT_DEPARTMENTAL int = 9
const STAMP_ANNOT_FORCOMMENT int = 10
const STAMP_ANNOT_TOPSECRET int = 11
const STAMP_ANNOT_DRAFT int = 12
const STAMP_ANNOT_FORPUBLICRELEASE int = 13

func (page *Page) Page_CreateStampAnnot(rect Rect, name int, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateStampAnnot(page.ptr, rect.toHaru(), C.HPDF_StampAnnotName(name), _text, encoder.ptr)
	return Annotation{ptr: rc}

}

func (page *Page) Page_CreateProjectionAnnot(rect Rect, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateProjectionAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)
	return Annotation{ptr: rc}
}

func (page *Page) Page_CreateSquareAnnot(rect Rect, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateSquareAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)
	return Annotation{ptr: rc}
}

func (page *Page) Page_CreateCircleAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateCircleAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

const ANNOT_NO_HIGHTLIGHT int = 0
const HPDF_ANNOT_INVERT_BOX int = 1
const HPDF_ANNOT_INVERT_BORDER int = 2
const HPDF_ANNOT_DOWN_APPEARANCE int = 3

func (annot *Annotation) LinkAnnot_SetHighlightMode(mode int) {
	C.HPDF_LinkAnnot_SetHighlightMode(annot.ptr, C.HPDF_AnnotHighlightMode(mode))
}

func (annot *Annotation) LinkAnnot_SetJavaScript(javascript JavaScript) {
	C.HPDF_LinkAnnot_SetJavaScript(annot.ptr, javascript.ptr)
}

func (annot *Annotation) LinkAnnot_SetBorderStyle(width float32, dash_on uint16, dash_off uint16) {
	C.HPDF_LinkAnnot_SetBorderStyle(annot.ptr, C.float(width), C.ushort(dash_on), C.ushort(dash_off))
}

const ANNOT_ICON_COMMENT int = 0
const ANNOT_ICON_KEY int = 1
const ANNOT_ICON_NOTE int = 2
const ANNOT_ICON_HELP int = 3
const ANNOT_ICON_NEW_PARAGRAPH int = 4
const ANNOT_ICON_PARAGRAPH int = 5
const ANNOT_ICON_INSERT int = 6

func (annot *Annotation) TextAnnot_SetIcon(icon int) {
	C.HPDF_TextAnnot_SetIcon(annot.ptr, C.HPDF_AnnotIcon(icon))
}

func (annot *Annotation) TextAnnot_SetOpened(opened bool) {
	var _opened C.int
	if opened {
		_opened = C.HPDF_TRUE
	} else {
		_opened = C.HPDF_FALSE
	}
	C.HPDF_TextAnnot_SetOpened(annot.ptr, _opened)
}

func (annot *Annotation) Annot_SetRGBColor(color RGBColor) {
	C.HPDF_Annot_SetRGBColor(annot.ptr, color.toHaru())
}

func (annot *Annotation) Annot_SetCMYKColor(color CMYKColor) {
	C.HPDF_Annot_SetCMYKColor(annot.ptr, color.toHaru())
}

func (annot *Annotation) Annot_SetGrayColor(color float32) {
	C.HPDF_Annot_SetGrayColor(annot.ptr, C.float(color))
}

func (annot *Annotation) Annot_SetNoColor() {
	C.HPDF_Annot_SetNoColor(annot.ptr)
}

func (annot *Annotation) MarkupAnnot_SetTitle(name string) {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.HPDF_MarkupAnnot_SetTitle(annot.ptr, _name)
}

func (annot *Annotation) MarkupAnnot_SetSubject(name string) {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.HPDF_MarkupAnnot_SetSubject(annot.ptr, _name)
}

func (annot *Annotation) MarkupAnnot_SetCreationDate(value time.Time) {
	C.HPDF_MarkupAnnot_SetCreationDate(annot.ptr, timeToHaru(value))
}

func (annot *Annotation) MarkupAnnot_SetTransparency(value float32) {
	C.HPDF_MarkupAnnot_SetTransparency(annot.ptr, C.float(value))
}

const ANNOT_INTENT_FREETEXTCALLOUT int = 0
const ANNOT_INTENT_FREETEXTTYPEWRITER int = 1
const ANNOT_INTENT_LINEARROW int = 2
const ANNOT_INTENT_LINEDIMENSION int = 3
const ANNOT_INTENT_POLYGONCLOUD int = 4
const ANNOT_INTENT_POLYLINEDIMENSION int = 5
const ANNOT_INTENT_POLYGONDIMENSION int = 6

func (annot *Annotation) MarkupAnnot_SetIntent(intent int) {
	C.HPDF_MarkupAnnot_SetIntent(annot.ptr, C.HPDF_AnnotIntent(intent))
}

func (annot *Annotation) MarkupAnnot_SetPopup(popup Annotation) {
	C.HPDF_MarkupAnnot_SetPopup(annot.ptr, popup.ptr)
}

func (annot *Annotation) MarkupAnnot_SetRectDiff(rect Rect) {
	C.HPDF_MarkupAnnot_SetRectDiff(annot.ptr, rect.toHaru())
}

func (annot *Annotation) MarkupAnnot_SetCloudEffect(cloudIntensity int) {
	C.HPDF_MarkupAnnot_SetCloudEffect(annot.ptr, C.int(cloudIntensity))
}

func (annot *Annotation) MarkupAnnot_SetInteriorRGBColor(color RGBColor) {
	C.HPDF_MarkupAnnot_SetInteriorRGBColor(annot.ptr, color.toHaru())
}

func (annot *Annotation) MarkupAnnot_SetInteriorCMYKColor(color CMYKColor) {
	C.HPDF_MarkupAnnot_SetInteriorCMYKColor(annot.ptr, color.toHaru())
}

func (annot *Annotation) MarkupAnnot_SetInteriorGrayColor(color float32) {
	C.HPDF_MarkupAnnot_SetInteriorGrayColor(annot.ptr, C.float(color))
}

func (annot *Annotation) MarkupAnnot_SetInteriorTransparent() {
	C.HPDF_MarkupAnnot_SetInteriorTransparent(annot.ptr)
}

func (annot *Annotation) TextMarkupAnnot_SetQuadPoints(lb Point, rb Point, rt Point, lt Point) {
	C.HPDF_TextMarkupAnnot_SetQuadPoints(annot.ptr, lb.toHaru(), rb.toHaru(), rt.toHaru(), lt.toHaru())
	return

}

func (annot *Annotation) PopupAnnot_SetOpened(opened bool) {
	var _opened C.int
	if opened {
		_opened = C.HPDF_TRUE
	} else {
		_opened = C.HPDF_FALSE
	}
	C.HPDF_PopupAnnot_SetOpened(annot.ptr, _opened)
}

const LINE_ANNOT_NONE int = 0
const LINE_ANNOT_SQUARE int = 1
const LINE_ANNOT_CIRCLE int = 2
const LINE_ANNOT_DIAMOND int = 3
const LINE_ANNOT_OPENARROW int = 4
const LINE_ANNOT_CLOSEDARROW int = 5
const LINE_ANNOT_BUTT int = 6
const LINE_ANNOT_ROPENARROW int = 7
const LINE_ANNOT_RCLOSEDARROW int = 8
const LINE_ANNOT_SLASH int = 9

func (annot *Annotation) FreeTextAnnot_SetLineEndingStyle(startStyle int, endStyle int) {
	C.HPDF_FreeTextAnnot_SetLineEndingStyle(annot.ptr, C.HPDF_LineAnnotEndingStyle(startStyle), C.HPDF_LineAnnotEndingStyle(endStyle))
}

func (annot *Annotation) FreeTextAnnot_Set3PointCalloutLine(startPoint Point, kneePoint Point, endPoint Point) {
	C.HPDF_FreeTextAnnot_Set3PointCalloutLine(annot.ptr, startPoint.toHaru(), kneePoint.toHaru(), endPoint.toHaru())
}

func (annot *Annotation) FreeTextAnnot_Set2PointCalloutLine(startPoint Point, endPoint Point) {
	C.HPDF_FreeTextAnnot_Set2PointCalloutLine(annot.ptr, startPoint.toHaru(), endPoint.toHaru())
}

func (annot *Annotation) FreeTextAnnot_SetDefaultStyle(style string) {
	_style := C.CString(style)
	defer C.free(unsafe.Pointer(_style))
	C.HPDF_FreeTextAnnot_SetDefaultStyle(annot.ptr, _style)
}

func (annot *Annotation) LineAnnot_SetPosition(startPoint Point, startStyle int, endPoint Point, endStyle int) {
	C.HPDF_LineAnnot_SetPosition(annot.ptr, startPoint.toHaru(), C.HPDF_LineAnnotEndingStyle(startStyle), endPoint.toHaru(), C.HPDF_LineAnnotEndingStyle(endStyle))
}

func (annot *Annotation) LineAnnot_SetLeader(leaderLen int, leaderExtLen int, leaderOffsetLen int) {
	C.HPDF_LineAnnot_SetLeader(annot.ptr, C.int(leaderLen), C.int(leaderExtLen), C.int(leaderOffsetLen))
}

const LINE_ANNOT_CAP_INLINE int = 0
const LINE_ANNOT_CAP_TOP int = 1

func (annot *Annotation) LineAnnot_SetCaption(showCaption bool, position int, horzOffset int, vertOffset int) {
	var _showCaption C.int
	if showCaption {
		_showCaption = C.HPDF_TRUE
	} else {
		_showCaption = C.HPDF_FALSE
	}
	C.HPDF_LineAnnot_SetCaption(annot.ptr, _showCaption, C.HPDF_LineAnnotCapPosition(position), C.int(horzOffset), C.int(vertOffset))
}

//subtype BS_*
func (annot *Annotation) Annotation_SetBorderStyle(subtype int, width float32, dash_on uint16, dash_off uint16, dash_phase uint16) {
	C.HPDF_Annotation_SetBorderStyle(annot.ptr, C.HPDF_BSSubtype(subtype), C.float(width), C.ushort(dash_on), C.ushort(dash_off), C.ushort(dash_phase))
}

func (annot *Annotation) ProjectionAnnot_SetExData(exdata ExData) {
	C.HPDF_ProjectionAnnot_SetExData(annot.ptr, exdata.ptr)
}
