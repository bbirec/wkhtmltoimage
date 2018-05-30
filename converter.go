package image

/*
#cgo LDFLAGS: -lwkhtmltox
#include <stdio.h>
#include <stdlib.h>
#include <wkhtmltox/image.h>

void errorCallback(wkhtmltoimage_converter * c, const char * msg) {
	fprintf(stderr, "Error: %s\n", msg);
}

typedef void (*err)(wkhtmltoimage_converter * c, const char * msg);
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Converter struct {
	converter *C.wkhtmltoimage_converter
	settings  *C.wkhtmltoimage_global_settings
}

func setOption(setting *C.wkhtmltoimage_global_settings, name string, value string) error {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	v := C.CString(value)
	defer C.free(unsafe.Pointer(v))

	if C.wkhtmltoimage_set_global_setting(setting, n, v) != 1 {
		return errors.New("Could not set option")
	}

	return nil
}

func NewConverter(opt map[string]string) (*Converter, error) {
	settings := C.wkhtmltoimage_create_global_settings()

	for key, value := range opt {
		if err := setOption(settings, key, value); err != nil {
			return nil, err
		}
	}

	converter := C.wkhtmltoimage_create_converter(settings, nil)

	return &Converter{converter: converter, settings: settings}, nil
}

func (c *Converter) Convert() ([]byte, error) {
	C.wkhtmltoimage_set_error_callback(c.converter, C.err(C.errorCallback))
	C.wkhtmltoimage_set_warning_callback(c.converter, C.err(C.errorCallback))

	if C.wkhtmltoimage_convert(c.converter) != 1 {
		return nil, errors.New("Conversion failed")
	}

	var output *C.uchar
	size := C.wkhtmltoimage_get_output(c.converter, &output)
	if size == 0 {
		return nil, errors.New("Could not retrieve converted object")
	}

	return C.GoBytes(unsafe.Pointer(output), C.int(size)), nil
}

func (c *Converter) Destroy() {
	C.wkhtmltoimage_destroy_converter(c.converter)
}
