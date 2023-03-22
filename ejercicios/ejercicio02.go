package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func TablaNumero() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese el número :")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())

			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
	}

}
