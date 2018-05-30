package image

/*
#cgo LDFLAGS: -lwkhtmltox
#include <stdlib.h>
#include <wkhtmltox/image.h>
*/
import "C"

func Init() {
	C.wkhtmltoimage_init(0)
}

func Version() string {
	return C.GoString(C.wkhtmltoimage_version())
}

func Destroy() {
	C.wkhtmltoimage_deinit()
}
