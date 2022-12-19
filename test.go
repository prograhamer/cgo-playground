package main

import (
	"fmt"
	"math/rand"

	"github.com/prograhamer/heap/internal/clib"
)

func main() {
	for i := 0; i < 10; i++ {
		numbers := randSlice(10)
		fmt.Println("numbers", i, numbers)

		tree, err := clib.NewTree()
		if err != nil {
			panic(err)
		}

		err = tree.Add(numbers...)
		if err != nil {
			panic(err)
		}

		err = tree.Walk()
		if err != nil {
			panic(err)
		}

		sorted, err := tree.Sort()
		fmt.Println("sorted", i, sorted)

		sorted = make([]int32, tree.Size()/2)
		tree.SortWithBuf(sorted)
		fmt.Println("sorted", i, sorted)

		clib.Destroy(tree)
	}
}

func randSlice(n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = rand.Intn(256)
	}

	return result
}
