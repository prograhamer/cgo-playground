package main

import (
	"github.com/prograhamer/heap/internal/clib"
)

func main() {
	tree, err := clib.NewTree()
	if err != nil {
		panic(err)
	}

	numbers := []int{10, 5, 7, 0, 15, 12, 2, 20, 9}

	err = tree.Add(numbers...)
	if err != nil {
		panic(err)
	}

	tree.Walk()
}
