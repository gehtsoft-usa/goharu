package goharu

//#cgo CFLAGS: -I${SRCDIR}/libharu/include
//#include <hpdf.h>
import "C"
import "unsafe"

//Image struct keeps data about an image
type Image struct {
	ptr C.HPDF_Image
}

//LoadPngImageFromFile loads an external PNG image file.
func (v *Doc) LoadPngImageFromFile(name string) Image {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	image := C.HPDF_LoadPngImageFromFile(v.ptr, _name)
	return Image{ptr: image}
}

//LoadPngImageFromMem loads PNG image from a buffer.
func (v *Doc) LoadPngImageFromMem(raw []byte) Image {
	image := C.HPDF_LoadPngImageFromMem(v.ptr, (*C.uchar)(unsafe.Pointer(&raw[0])), C.uint(len(raw)))
	return Image{ptr: image}
}

//LoadPngImageFromFile2 loads a PNG image from the file
//
//The image isn't actually loaded until it is used in the document
//LoadPngImageFromFile2 loads an external PNG image file.
func (v *Doc) LoadPngImageFromFile2(name string) Image {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	image := C.HPDF_LoadPngImageFromFile2(v.ptr, _name)
	return Image{ptr: image}
}

//LoadJpegImageFromFile loads an external JPEG image file.
func (v *Doc) LoadJpegImageFromFile(name string) Image {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	image := C.HPDF_LoadJpegImageFromFile(v.ptr, _name)
	return Image{ptr: image}
}

//LoadJpegImageFromMem loads JPEG image from a buffer.
func (v *Doc) LoadJpegImageFromMem(raw []byte) Image {
	image := C.HPDF_LoadJpegImageFromMem(v.ptr, (*C.uchar)(unsafe.Pointer(&raw[0])), C.uint(len(raw)))
	return Image{ptr: image}
}

//Size returns width and height of the image
func (i *Image) Size() (uint, uint) {
	return uint(C.HPDF_Image_GetWidth(i.ptr)), uint(C.HPDF_Image_GetHeight(i.ptr))
}

//Width returns width of the image
func (i *Image) Width() uint {
	return uint(C.HPDF_Image_GetWidth(i.ptr))
}

//Height returns height of the image
func (i *Image) Height() uint {
	return uint(C.HPDF_Image_GetHeight(i.ptr))
}

//BitsPerComponent returns bits per pixel of the image
func (i *Image) BitsPerComponent() uint {
	return uint(C.HPDF_Image_GetBitsPerComponent(i.ptr))
}

//SetColorMask sets mask using color components
func (i *Image) SetColorMask(redmin, redmax, greenmin, greenmax, bluemin, bluemax uint) {
	C.HPDF_Image_SetColorMask(i.ptr, C.uint(redmin), C.uint(redmax), C.uint(greenmin), C.uint(greenmax), C.uint(bluemin), C.uint(bluemax))
}

//SetImageMask set image mask using another image as a mask
func (i *Image) SetImageMask(mask *Image) {
	C.HPDF_Image_SetMaskImage(i.ptr, mask.ptr)
}

//AddSMask sets mark using another image
func (i *Image) AddSMask(mask *Image) {
	C.HPDF_Image_AddSMask(i.ptr, mask.ptr)
}

//GetColorSpace gets color space of the image
func (i *Image) GetColorSpace() string {
	return C.GoString(C.HPDF_Image_GetColorSpace(i.ptr))
}
