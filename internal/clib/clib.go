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

func Destroy(tree *Tree) error {
	_, err := C.HETreeFree(tree.ctree)
	if err != nil {
		return fmt.Errorf("HETreeFree: %w", err)
	}
	return nil
}

func (t *Tree) Size() int {
	return int(t.ctree.size)
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

func (t *Tree) Sort() ([]int32, error) {
	res, err := C.HETreeSort(t.ctree)
	if err != nil {
		return nil, fmt.Errorf("HETreeSort: %w", err)
	}
	defer C.free(unsafe.Pointer(res))

	sorted := make([]int32, t.ctree.size)
	for i := int32(0); i < int32(t.ctree.size); i++ {
		sorted[i] = *(*int32)(unsafe.Add(unsafe.Pointer(res), i*4))
	}

	return sorted, nil
}

func (t *Tree) Sorted() ([]int32, error) {
	res := make([]int32, int(t.ctree.size))
	_, err := C.HETreeSortNoMalloc(t.ctree, (*C.int)(&res[0]), t.ctree.size)
	if err != nil {
		return nil, fmt.Errorf("HETreeSortNoMalloc: %w", err)
	}
	return res, nil
}

func (t *Tree) SortWithBuf(buf []int32) error {
	fmt.Println("cap:", cap(buf), "len:", len(buf))
	ccap := C.int(cap(buf))
	cres := &buf[0]

	count, err := C.HETreeSortNoMalloc(t.ctree, (*C.int)(cres), ccap)
	if err != nil {
		return fmt.Errorf("HETreeSortNoMalloc: %w", err)
	}

	fmt.Println("count:", count)

	return nil
}
