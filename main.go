package main

import (
	"fmt"

	"github.com/alvarez-alberto/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(123)

	fmt.Println(estado)
	fmt.Println(texto)
}
