package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		estado, texto := variables.ConviertoaTexto(123)

		fmt.Println(estado)
		fmt.Println(texto)
	*/
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("No es windows")
	} else {
		fmt.Println("windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux.":
		fmt.Println("Esto es Linux")
	case "Darwin.":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
