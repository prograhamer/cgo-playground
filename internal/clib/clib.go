package clib

//#include "clib.h"
//#include <stdlib.h>
import "C"
import "unsafe"

func Reverse(data []byte) ([]byte, error) {
	clen := C.ulong(len(data))
	cdata := C.CString(string(data))
	defer C.free(unsafe.Pointer(cdata))

	crev, err := C.Reverse(cdata, clen)
	if err != nil {
		return nil, err
	}
	defer C.free(unsafe.Pointer(crev))

	res := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = *(*byte)(unsafe.Add(unsafe.Pointer(crev), i))
	}

	return res, nil
}

func ReverseInPlace(data []byte) error {
	clen := C.ulong(len(data))
	cdata := C.CString(string(data))
	defer C.free(unsafe.Pointer(cdata))

	_, err := C.ReverseInPlace(cdata, clen)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i++ {
		data[i] = *(*byte)(unsafe.Add(unsafe.Pointer(cdata), i))
	}
	return nil
}
