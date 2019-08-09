package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import (
	"time"
	"unsafe"
)

//Annotation struct refers to haru annotation object
type Annotation struct {
	ptr C.HPDF_Annotation
}

//Create3DAnnot creates an annotation that contains a 3D object
func (page *Page) Create3DAnnot(rect Rect, tb bool, np bool, u3d U3D, ap Image) Annotation {
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

//CreateTextAnnot creates a new text annotation object for the page.
func (page *Page) CreateTextAnnot(rect Rect, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateTextAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)
	return Annotation{ptr: rc}
}

//CreateFreeTextAnnot creates free text annotation
func (page *Page) CreateFreeTextAnnot(rect Rect, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateFreeTextAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)
	return Annotation{ptr: rc}
}

//CreateLineAnnot creates line annotation
func (page *Page) CreateLineAnnot(text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateLineAnnot(page.ptr, _text, encoder.ptr)
	return Annotation{ptr: rc}
}

//CreateWidgetAnnotWhiteOnlyWhilePrint creates widget annotation that exists only at prints
func (pdf *Doc) CreateWidgetAnnotWhiteOnlyWhilePrint(page Page, rect Rect) Annotation {

	rc := C.HPDF_Page_CreateWidgetAnnot_WhiteOnlyWhilePrint(pdf.ptr, page.ptr, rect.toHaru())
	return Annotation{ptr: rc}
}

//CreateWidgetAnnot creates widget-based annotation
func (page *Page) CreateWidgetAnnot(rect Rect) Annotation {

	rc := C.HPDF_Page_CreateWidgetAnnot(page.ptr, rect.toHaru())

	return Annotation{ptr: rc}

}

//CreateLinkAnnot creates a new link annotation object for the page.
func (page *Page) CreateLinkAnnot(rect Rect, dst Destination) Annotation {
	rc := C.HPDF_Page_CreateLinkAnnot(page.ptr, rect.toHaru(), dst.ptr)
	return Annotation{ptr: rc}

}

//CreateURILinkAnnot creates a new web link annotation object for the page.
func (page *Page) CreateURILinkAnnot(rect Rect, uri string) Annotation {
	_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(_uri))
	rc := C.HPDF_Page_CreateURILinkAnnot(page.ptr, rect.toHaru(), _uri)
	return Annotation{ptr: rc}

}

//AnnotTextNotes is value that indicates annotation type
const AnnotTextNotes int = 0

//AnnotLink is value  that indicates annotation type
const AnnotLink int = 1

//AnnotSound is value  that indicates annotation type
const AnnotSound int = 2

//AnnotFreeText is value  that indicates annotation type
const AnnotFreeText int = 3

//AnnotStamp is value  that indicates annotation type
const AnnotStamp int = 4

//AnnotSquare is value  that indicates annotation type
const AnnotSquare int = 5

//AnnotCircle is value  that indicates annotation type
const AnnotCircle int = 6

//AnnotStrikeOut is value  that indicates annotation type
const AnnotStrikeOut int = 7

//AnnotHightlight is value  that indicates annotation type
const AnnotHightlight int = 8

//AnnotUnderline is value  that indicates annotation type
const AnnotUnderline int = 9

//AnnotInk is value  that indicates annotation type
const AnnotInk int = 10

//AnnotFileAttachment is value  that indicates annotation type
const AnnotFileAttachment int = 11

//AnnotPopup is value  that indicates annotation type
const AnnotPopup int = 12

//Annot3d is value  that indicates annotation type
const Annot3d int = 13

//AnnotSquiggly is value  that indicates annotation type
const AnnotSquiggly int = 14

//AnnotLine is value  that indicates annotation type
const AnnotLine int = 15

//AnnotProjection is value  that indicates annotation type
const AnnotProjection int = 16

//AnnotWidget is value  that indicates annotation type
const AnnotWidget int = 17

//CreateTextMarkupAnnot creates markup annotation
func (page *Page) CreateTextMarkupAnnot(rect Rect, text string, encoder Encoder, subType int) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateTextMarkupAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr, C.HPDF_AnnotType(subType))
	return Annotation{ptr: rc}

}

//CreateHighlightAnnot creates highlighted annotation
func (page *Page) CreateHighlightAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateHighlightAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

//CreateUnderlineAnnot creates underlined annotation
func (page *Page) CreateUnderlineAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateUnderlineAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

//CreateSquigglyAnnot creates squiggly annotation
func (page *Page) CreateSquigglyAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateSquigglyAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

//CreateStrikeOutAnnot creates stoke out annotation
func (page *Page) CreateStrikeOutAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateStrikeOutAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

//CreatePopupAnnot creates a popup annotation
func (page *Page) CreatePopupAnnot(rect Rect, parent Annotation) Annotation {

	rc := C.HPDF_Page_CreatePopupAnnot(page.ptr, rect.toHaru(), parent.ptr)

	return Annotation{ptr: rc}

}

//StampAnnotApproved is value indicating "APPROVED" stamp type
const StampAnnotApproved int = 0

//StampAnnotExperimental is value indicating "EXPERIMENTAL" stamp type
const StampAnnotExperimental int = 1

//StampAnnotNotapproved is value indicating  "NOT APPROVED" stamp type
const StampAnnotNotapproved int = 2

//StampAnnotAsis is value indicating  "AS IS" stamp type
const StampAnnotAsis int = 3

//StampAnnotExpired is value indicating "EXPIRED" stamp type
const StampAnnotExpired int = 4

//StampAnnotNotForPublicRelease is value indicating "NOT FOR PUBLIC RELEASE" stamp type
const StampAnnotNotForPublicRelease int = 5

//StampAnnotConfidential is value indicating "CONFIDENTIAL stamp type"
const StampAnnotConfidential int = 6

//StampAnnotFinal is value indicating "FINAL" stamp type
const StampAnnotFinal int = 7

//StampAnnotSold is value indicating "SOLD" stamp type
const StampAnnotSold int = 8

//StampAnnotDepartmental is value indicating "DEPARTMENTAL" stamp type
const StampAnnotDepartmental int = 9

//StampAnnotForComment is value indicating "FOR COMMENT" stamp type
const StampAnnotForComment int = 10

//StampAnnotTopSecret is value indicating "TOP SECRET" stamp type
const StampAnnotTopSecret int = 11

//StampAnnotDraft is value indicating "DRAFT" stamp type
const StampAnnotDraft int = 12

//StampAnnotForPublicRelease is value indicating "FOR PUBLIC RELEASE" stamp type
const StampAnnotForPublicRelease int = 13

//CreateStampAnnot creates stamp annotation.
//
//Use StampAnnot* values to specify the type of the stamp
func (page *Page) CreateStampAnnot(rect Rect, name int, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateStampAnnot(page.ptr, rect.toHaru(), C.HPDF_StampAnnotName(name), _text, encoder.ptr)
	return Annotation{ptr: rc}

}

//CreateProjectionAnnot creates projection annotation
func (page *Page) CreateProjectionAnnot(rect Rect, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateProjectionAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)
	return Annotation{ptr: rc}
}

//CreateSquareAnnot creates square annotation
func (page *Page) CreateSquareAnnot(rect Rect, text string, encoder Encoder) Annotation {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateSquareAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)
	return Annotation{ptr: rc}
}

//CreateCircleAnnot creates circle annotation
func (page *Page) CreateCircleAnnot(rect Rect, text string, encoder Encoder) Annotation {

	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	rc := C.HPDF_Page_CreateCircleAnnot(page.ptr, rect.toHaru(), _text, encoder.ptr)

	return Annotation{ptr: rc}

}

//AnnotNoHightlight is highlight mode: no highlight
const AnnotNoHightlight int = 0

//AnnotInvertBox is highlight mode: invert content
const AnnotInvertBox int = 1

//AnnotInvertBorder is highlight mode: invert border
const AnnotInvertBorder int = 2

//AnnotDownAppearance is highlight mode:
const AnnotDownAppearance int = 3

//LinkAnnotSetHighlightMode defines the appearance when a mouse clicks on a link annotation. (see example program)
//
//Use AnnotNoHightlight, AnnotInvertBox, AnnotInvertBorder or AnnotDownAppearance as mode
func (annot *Annotation) LinkAnnotSetHighlightMode(mode int) {
	C.HPDF_LinkAnnot_SetHighlightMode(annot.ptr, C.HPDF_AnnotHighlightMode(mode))
}

//LinkAnnotSetJavaScript sets javascript to executed by the link annotation
func (annot *Annotation) LinkAnnotSetJavaScript(javascript JavaScript) {
	C.HPDF_LinkAnnot_SetJavaScript(annot.ptr, javascript.ptr)
}

//LinkAnnotSetBorderStyle defines the style of the annotation's border.
func (annot *Annotation) LinkAnnotSetBorderStyle(width float32, dashOn uint16, dashOff uint16) {
	C.HPDF_LinkAnnot_SetBorderStyle(annot.ptr, C.float(width), C.ushort(dashOn), C.ushort(dashOff))
}

//AnnotIconComment is icon type for icon annotation
const AnnotIconComment int = 0

//AnnotIconKey is icon type for icon annotation
const AnnotIconKey int = 1

//AnnotIconNote is icon type for icon annotation
const AnnotIconNote int = 2

//AnnotIconHelp is icon type for icon annotation
const AnnotIconHelp int = 3

//AnnotIconNewParagraph is icon type for icon annotation
const AnnotIconNewParagraph int = 4

//AnnotIconParagraph is icon type for icon annotation
const AnnotIconParagraph int = 5

//AnnotIconInsert is icon type for icon annotation
const AnnotIconInsert int = 6

//TextAnnotSetIcon defines the style of the annotation's icon. (see example program)
//
//Use AnnotIcon* values as icon parameter
func (annot *Annotation) TextAnnotSetIcon(icon int) {
	C.HPDF_TextAnnot_SetIcon(annot.ptr, C.HPDF_AnnotIcon(icon))
}

//TextAnnotSetOpened defines whether the text-annotation is initially open.
func (annot *Annotation) TextAnnotSetOpened(opened bool) {
	var _opened C.int
	if opened {
		_opened = C.HPDF_TRUE
	} else {
		_opened = C.HPDF_FALSE
	}
	C.HPDF_TextAnnot_SetOpened(annot.ptr, _opened)
}

//AnnotSetRGBColor sets annotation color
func (annot *Annotation) AnnotSetRGBColor(color RGBColor) {
	C.HPDF_Annot_SetRGBColor(annot.ptr, color.toHaru())
}

//AnnotSetCMYKColor set annotation color
func (annot *Annotation) AnnotSetCMYKColor(color CMYKColor) {
	C.HPDF_Annot_SetCMYKColor(annot.ptr, color.toHaru())
}

//AnnotSetGrayColor sets annotation color as a shade of gray
func (annot *Annotation) AnnotSetGrayColor(color float32) {
	C.HPDF_Annot_SetGrayColor(annot.ptr, C.float(color))
}

//AnnotSetNoColor removes annotation color information
func (annot *Annotation) AnnotSetNoColor() {
	C.HPDF_Annot_SetNoColor(annot.ptr)
}

//MarkupAnnotSetTitle sets markup annotation title
func (annot *Annotation) MarkupAnnotSetTitle(name string) {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.HPDF_MarkupAnnot_SetTitle(annot.ptr, _name)
}

//MarkupAnnotSetSubject sets markup annotation subject
func (annot *Annotation) MarkupAnnotSetSubject(name string) {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	C.HPDF_MarkupAnnot_SetSubject(annot.ptr, _name)
}

//MarkupAnnotSetCreationDate sets markup annotation date
func (annot *Annotation) MarkupAnnotSetCreationDate(value time.Time) {
	C.HPDF_MarkupAnnot_SetCreationDate(annot.ptr, timeToHaru(value))
}

//MarkupAnnotSetTransparency sets markup annotation transparency
func (annot *Annotation) MarkupAnnotSetTransparency(value float32) {
	C.HPDF_MarkupAnnot_SetTransparency(annot.ptr, C.float(value))
}

//AnnotIntentFreetextCallout is markup annotation intent
const AnnotIntentFreetextCallout int = 0

//AnnotIntentFreetexttypewriter is markup annotation intent
const AnnotIntentFreetexttypewriter int = 1

//AnnotIntentLinearrow is markup annotation intent
const AnnotIntentLinearrow int = 2

//AnnotIntentLinedimension is markup annotation intent
const AnnotIntentLinedimension int = 3

//AnnotIntentPolygoncloud is markup annotation intent
const AnnotIntentPolygoncloud int = 4

//AnnotIntentPolylinedimension is markup annotation intent
const AnnotIntentPolylinedimension int = 5

//AnnotIntentPolygondimension is markup annotation intent
const AnnotIntentPolygondimension int = 6

//MarkupAnnotSetIntent sets intent of markup annotation
func (annot *Annotation) MarkupAnnotSetIntent(intent int) {
	C.HPDF_MarkupAnnot_SetIntent(annot.ptr, C.HPDF_AnnotIntent(intent))
}

//MarkupAnnotSetPopup sets annotation that popups at markup annotation
func (annot *Annotation) MarkupAnnotSetPopup(popup Annotation) {
	C.HPDF_MarkupAnnot_SetPopup(annot.ptr, popup.ptr)
}

//MarkupAnnotSetRectDiff sets markup annotation rectangle
func (annot *Annotation) MarkupAnnotSetRectDiff(rect Rect) {
	C.HPDF_MarkupAnnot_SetRectDiff(annot.ptr, rect.toHaru())
}

//MarkupAnnotSetCloudEffect sets markup cloud effect intensity
func (annot *Annotation) MarkupAnnotSetCloudEffect(cloudIntensity int) {
	C.HPDF_MarkupAnnot_SetCloudEffect(annot.ptr, C.int(cloudIntensity))
}

//MarkupAnnotSetInteriorRGBColor sets markup annotation interior color
func (annot *Annotation) MarkupAnnotSetInteriorRGBColor(color RGBColor) {
	C.HPDF_MarkupAnnot_SetInteriorRGBColor(annot.ptr, color.toHaru())
}

//MarkupAnnotSetInteriorCMYKColor sets markup annotation interior color
func (annot *Annotation) MarkupAnnotSetInteriorCMYKColor(color CMYKColor) {
	C.HPDF_MarkupAnnot_SetInteriorCMYKColor(annot.ptr, color.toHaru())
}

//MarkupAnnotSetInteriorGrayColor sets markup annotation interior color as a shade of gray
func (annot *Annotation) MarkupAnnotSetInteriorGrayColor(color float32) {
	C.HPDF_MarkupAnnot_SetInteriorGrayColor(annot.ptr, C.float(color))
}

//MarkupAnnotSetInteriorTransparent makes markup annotation interior transparent
func (annot *Annotation) MarkupAnnotSetInteriorTransparent() {
	C.HPDF_MarkupAnnot_SetInteriorTransparent(annot.ptr)
}

//TextMarkupAnnotSetQuadPoints sets markup annotation coordinates
func (annot *Annotation) TextMarkupAnnotSetQuadPoints(lb Point, rb Point, rt Point, lt Point) {
	C.HPDF_TextMarkupAnnot_SetQuadPoints(annot.ptr, lb.toHaru(), rb.toHaru(), rt.toHaru(), lt.toHaru())
	return

}

//PopupAnnotSetOpened makes popup annotation opened or closed by default
func (annot *Annotation) PopupAnnotSetOpened(opened bool) {
	var _opened C.int
	if opened {
		_opened = C.HPDF_TRUE
	} else {
		_opened = C.HPDF_FALSE
	}
	C.HPDF_PopupAnnot_SetOpened(annot.ptr, _opened)
}

//LineAnnotNone is line annotation type
const LineAnnotNone int = 0

//LineAnnotSquare is line annotation type
const LineAnnotSquare int = 1

//LineAnnotCircle is line annotation type
const LineAnnotCircle int = 2

//LineAnnotDiamond is line annotation type
const LineAnnotDiamond int = 3

//LineAnnotOpenarrow is line annotation type
const LineAnnotOpenarrow int = 4

//LineAnnotClosedarrow is line annotation type
const LineAnnotClosedarrow int = 5

//LineAnnotButt is line annotation type
const LineAnnotButt int = 6

//LineAnnotROpenArrow is line annotation type
const LineAnnotROpenArrow int = 7

//LineAnnotRClosedArrow is line annotation type
const LineAnnotRClosedArrow int = 8

//LineAnnotSlash is line annotation type
const LineAnnotSlash int = 9

//FreeTextAnnotSetLineEndingStyle sets end line style for free text annotation
//
//Use LineAnnot* values to set style
func (annot *Annotation) FreeTextAnnotSetLineEndingStyle(startStyle int, endStyle int) {
	C.HPDF_FreeTextAnnot_SetLineEndingStyle(annot.ptr, C.HPDF_LineAnnotEndingStyle(startStyle), C.HPDF_LineAnnotEndingStyle(endStyle))
}

//FreeTextAnnotSet3PointCalloutLine sets free text annotation callout line coordinates
func (annot *Annotation) FreeTextAnnotSet3PointCalloutLine(startPoint Point, kneePoint Point, endPoint Point) {
	C.HPDF_FreeTextAnnot_Set3PointCalloutLine(annot.ptr, startPoint.toHaru(), kneePoint.toHaru(), endPoint.toHaru())
}

//FreeTextAnnotSet2PointCalloutLine set free text annotation callout line coordinates
func (annot *Annotation) FreeTextAnnotSet2PointCalloutLine(startPoint Point, endPoint Point) {
	C.HPDF_FreeTextAnnot_Set2PointCalloutLine(annot.ptr, startPoint.toHaru(), endPoint.toHaru())
}

//FreeTextAnnotSetDefaultStyle sets free text annotation style
func (annot *Annotation) FreeTextAnnotSetDefaultStyle(style string) {
	_style := C.CString(style)
	defer C.free(unsafe.Pointer(_style))
	C.HPDF_FreeTextAnnot_SetDefaultStyle(annot.ptr, _style)
}

//LineAnnotSetPosition sets line annotation position
func (annot *Annotation) LineAnnotSetPosition(startPoint Point, startStyle int, endPoint Point, endStyle int) {
	C.HPDF_LineAnnot_SetPosition(annot.ptr, startPoint.toHaru(), C.HPDF_LineAnnotEndingStyle(startStyle), endPoint.toHaru(), C.HPDF_LineAnnotEndingStyle(endStyle))
}

//LineAnnotSetLeader sets line annotation leader
func (annot *Annotation) LineAnnotSetLeader(leaderLen int, leaderExtLen int, leaderOffsetLen int) {
	C.HPDF_LineAnnot_SetLeader(annot.ptr, C.int(leaderLen), C.int(leaderExtLen), C.int(leaderOffsetLen))
}

//LineAnnotCapInline is line annotation line caption type
const LineAnnotCapInline int = 0

//LineAnnotCapTop is line annotation line caption type
const LineAnnotCapTop int = 1

//LineAnnotSetCaption sets line annotation line caption
//
//use LineAnnotationCap* values
func (annot *Annotation) LineAnnotSetCaption(showCaption bool, position int, horzOffset int, vertOffset int) {
	var _showCaption C.int
	if showCaption {
		_showCaption = C.HPDF_TRUE
	} else {
		_showCaption = C.HPDF_FALSE
	}
	C.HPDF_LineAnnot_SetCaption(annot.ptr, _showCaption, C.HPDF_LineAnnotCapPosition(position), C.int(horzOffset), C.int(vertOffset))
}

//SetBorderStyle defines the appearance of a text annotation.
//
//Use Bs* value as a subtype
func (annot *Annotation) SetBorderStyle(subtype int, width float32, dashOn uint16, dashOff uint16, dashPhase uint16) {
	C.HPDF_Annotation_SetBorderStyle(annot.ptr, C.HPDF_BSSubtype(subtype), C.float(width), C.ushort(dashOn), C.ushort(dashOff), C.ushort(dashPhase))
}

//SetExData associates the annotation with ExData object
func (annot *Annotation) SetExData(exdata ExData) {
	C.HPDF_ProjectionAnnot_SetExData(annot.ptr, exdata.ptr)
}
