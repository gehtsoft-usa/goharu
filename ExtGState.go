package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//ExtGState struct
type ExtGState struct {
	ptr C.HPDF_ExtGState
}

//CreateExtGState create G-state
func (pdf *Doc) CreateExtGState() ExtGState {
	return ExtGState{ptr: C.HPDF_CreateExtGState(pdf.ptr)}
}

//SetAlphaStroke sets alpha stroke
func (ext_gstate *ExtGState) SetAlphaStroke(value float32) {
	C.HPDF_ExtGState_SetAlphaStroke(ext_gstate.ptr, C.float(value))
}

//SetAlphaFill sets alpha fill
func (ext_gstate *ExtGState) SetAlphaFill(value float32) {
	C.HPDF_ExtGState_SetAlphaFill(ext_gstate.ptr, C.float(value))
}

//BmNormal is blending mode constant
const BmNormal int = 0 
//BmMultiply is blending mode constant
const BmMultiply int = 1 
//BmScreen is blending mode constant
const BmScreen int = 2 
//BmOverlay is blending mode constant
const BmOverlay int = 3 
//BmDarken is blending mode constant
const BmDarken int = 4 
//BmLighten is blending mode constant
const BmLighten int = 5 
//BmColorDodge is blending mode constant
const BmColorDodge int = 6 
//BmColorBum is blending mode constant
const BmColorBum int = 7 
//BmHardLight is blending mode constant
const BmHardLight int = 8 
//BmSoftLight is blending mode constant
const BmSoftLight int = 9 
//BmDifference is blending mode constant
const BmDifference int = 10 
//BmExclushon is blending mode constant
const BmExclushon int = 11 

//SetBlendMode sets blending mode
//
//Use Bm* value as the parameter
func (ext_gstate *ExtGState) SetBlendMode(mode int) {

	C.HPDF_ExtGState_SetBlendMode(ext_gstate.ptr, C.HPDF_BlendMode(mode))
}

//SetExtGState applies the graphics state to the page.
func (page *Page) SetExtGState(gstate ExtGState) {
	C.HPDF_Page_SetExtGState(page.ptr, gstate.ptr)
}
