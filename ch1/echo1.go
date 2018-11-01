package main

import (
	"fmt"
	"os"
)

func main() {
	joinArgs()
}

func joinArgs() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, os.Args[i])

		s += sep + os.Args[i]
		sep = " - "
	}

	fmt.Println(s)
}
