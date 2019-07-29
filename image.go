package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//The structure represents an image
type Image struct {
	ptr C.HPDF_Image
}

//Loads a PNG image from the file
func (v *Doc) LoadPngImageFromFile(name string) Image {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	image := C.HPDF_LoadPngImageFromFile(v.ptr, _name)
	return Image{ptr: image}
}

//Loads a PNG image from the memory
func (v *Doc) LoadPngImageFromMem(raw []byte) Image {
	image := C.HPDF_LoadPngImageFromMem(v.ptr, (*C.uchar)(unsafe.Pointer(&raw[0])), C.uint(len(raw)))
	return Image{ptr: image}
}

//Loads a PNG image from the file
//
//The image isn't actually loaded until it is used in the document
func (v *Doc) LoadPngImageFromFile2(name string) Image {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	image := C.HPDF_LoadPngImageFromFile2(v.ptr, _name)
	return Image{ptr: image}
}

//Loads a Jpeg image from the file
func (v *Doc) LoadJpegImageFromFile(name string) Image {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	image := C.HPDF_LoadJpegImageFromFile(v.ptr, _name)
	return Image{ptr: image}
}

//Loads a Jpeg image from the memory
func (v *Doc) LoadJpegImageFromMem(raw []byte) Image {
	image := C.HPDF_LoadJpegImageFromMem(v.ptr, (*C.uchar)(unsafe.Pointer(&raw[0])), C.uint(len(raw)))
	return Image{ptr: image}
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

//Returns bits per pixel of the image
func (i *Image) BitsPerComponent() uint {
	return uint(C.HPDF_Image_GetBitsPerComponent(i.ptr))
}

//Returns bits per pixel of the image
func (i *Image) SetColorMask(redmin, redmax, greenmin, greenmax, bluemin, bluemax uint) {
	C.HPDF_Image_SetColorMask(i.ptr, C.uint(redmin), C.uint(redmax), C.uint(greenmin), C.uint(greenmax), C.uint(bluemin), C.uint(bluemax))
}

func (i *Image) SetImageMask(mask *Image) {
	C.HPDF_Image_SetMaskImage(i.ptr, mask.ptr)
}
