package clib

//#include "clib.h"
//#include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Tree struct {
	ctree *C.HETree
}

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

func NewTree() (*Tree, error) {
	tree, err := C.HETreeInit()
	if err != nil {
		return nil, err
	}

	return &Tree{tree}, nil
}

func (t *Tree) Add(numbers ...int) error {
	for _, n := range numbers {
		_, err := C.HETreeAdd(t.ctree, C.int(n))
		if err != nil {
			return fmt.Errorf("HETreeAdd: %w", err)
		}
	}
	return nil
}

func (t *Tree) Walk() error {
	_, err := C.HETreePrint(t.ctree)
	if err != nil {
		return fmt.Errorf("HETreePrint: %w", err)
	}
	return nil
}
