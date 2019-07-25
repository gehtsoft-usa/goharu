package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"

type Point struct {
	X float32
	Y float32
}

//The structure represents an image
type Image struct {
	ptr C.HPDF_Image
}

//Returns width and height of the image
func (i *Image) Size() (uint, uint) {
	return uint(C.HPDF_Image_GetWidth(i.ptr)), uint(C.HPDF_Image_GetHeight(i.ptr))
}

//Returns width of the image
func (i *Image) Width() uint {
	return uint(C.HPDF_Image_GetWidth(i.ptr))
}

//Returns height of the image
func (i *Image) Height() uint {
	return uint(C.HPDF_Image_GetHeight(i.ptr))
}
