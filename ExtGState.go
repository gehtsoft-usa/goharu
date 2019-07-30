package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

//The structure represents an image
type ExtGState struct {
	ptr C.HPDF_ExtGState
}

func (pdf *Doc) CreateExtGState() ExtGState {
	return ExtGState{ptr: C.HPDF_CreateExtGState(pdf.ptr)}
}

func (ext_gstate *ExtGState) SetAlphaStroke(value float32) {
	C.HPDF_ExtGState_SetAlphaStroke(ext_gstate.ptr, C.float(value))
}

func (ext_gstate *ExtGState) SetAlphaFill(value float32) {
	C.HPDF_ExtGState_SetAlphaFill(ext_gstate.ptr, C.float(value))
}

const BM_NORMAL int = 0
const BM_MULTIPLY int = 1
const BM_SCREEN int = 2
const BM_OVERLAY int = 3
const BM_DARKEN int = 4
const BM_LIGHTEN int = 5
const BM_COLOR_DODGE int = 6
const BM_COLOR_BUM int = 7
const BM_HARD_LIGHT int = 8
const BM_SOFT_LIGHT int = 9
const BM_DIFFERENCE int = 10
const BM_EXCLUSHON int = 11
const BM_EOF int = 12

func (ext_gstate *ExtGState) SetBlendMode(mode int) {

	C.HPDF_ExtGState_SetBlendMode(ext_gstate.ptr, C.HPDF_BlendMode(mode))
}

func (page *Page) Page_SetExtGState(ext_gstate ExtGState) {
	C.HPDF_Page_SetExtGState(page.ptr, ext_gstate.ptr)
}
