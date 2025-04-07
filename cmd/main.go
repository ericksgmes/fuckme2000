package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fuckme2000 <file.f2000>")
		os.Exit(1)
	}

	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("Running program:")
	fmt.Println(string(src))
}