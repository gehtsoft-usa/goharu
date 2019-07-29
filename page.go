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

func (page *Page) Page_SetWidth(value float32) {
	C.HPDF_Page_SetWidth(page.ptr, C.float(value))
}

func (page *Page) Page_SetHeight(value float32) {
	C.HPDF_Page_SetHeight(page.ptr, C.float(value))
}

func (page *Page) Page_SetRotate(angle uint16) {
	C.HPDF_Page_SetRotate(page.ptr, C.ushort(angle))
}

func (page *Page) Page_SetZoom(zoom float32) {
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
