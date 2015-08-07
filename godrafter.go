package godrafter

// #cgo LDFLAGS: -lstdc++ -ldrafter -lsos -lsnowcrash -lmarkdownparser -lsundown
/*
#include <stdlib.h>
#include <string.h>
#include "cdrafter.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func DrafterParse(source []byte, flags int) ([]byte, error) {
	var err error
	var res unsafe.Pointer
	var e C.int
	e = C.drafter_c_parse(
		(*C.char)(unsafe.Pointer(&source[0])),
		C.sc_blueprint_parser_options(flags),
		(**C.char)(unsafe.Pointer(&res)),
	)
	if res == nil {
		return nil, fmt.Errorf("nil result pointer")
	}
	defer C.free(res)
	length := int64(C.strlen((*C.char)(res)))
	if int(e) != 0 {
		err = fmt.Errorf("error while parsing blueprint")
	}
	b := C.GoBytes(unsafe.Pointer(res), C.int(length))
	return b, err
}
