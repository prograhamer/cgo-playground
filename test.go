package main

import (
	"fmt"

	"github.com/prograhamer/heap/internal/clib"
)

func main() {
	message := []byte("Hello, world!")

	fmt.Println("message", message)

	reversed, err := clib.Reverse(message)
	if err != nil {
		panic(err)
	}

	fmt.Println("reversed", reversed)

	fmt.Println("message", message)
	clib.ReverseInPlace(message)
	fmt.Println("message", message)
}
