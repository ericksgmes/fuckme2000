package vm

import "fmt"

var customPrinter func(string) = nil

func SetPrinter(fn func(string)) {
	customPrinter = fn
}

func printlnVal(s string) {
	if customPrinter != nil {
		customPrinter(s)
	} else {
		fmt.Println(s)
	}
}
