package main

import (
	"fmt"
)

func main() {
	a, b := 21, 15
	var g, p, A int

	_, err := fmt.Scanf("g is %d and p is %d", &g, &p)

	if err != nil {
		fmt.Println("Error al leer los valores:", err)
	} else {
		fmt.Printf("OK\n")
	}

	_, err = fmt.Scanf("A is %d", &A)

	if err != nil {
		fmt.Println("Error al leer los valores:", err)
	}

	B := modularExponentiation(g, p, b)
	Sa := modularExponentiation(A, p, b)
	Sb := modularExponentiation(B, p, a)

	fmt.Printf("B is %d\n", B)
	fmt.Printf("A is %d\n", A)

	if Sa == Sb {
		fmt.Printf("S is %d", Sa)
	} else {
		panic("ALgo salio mal")
	}

}

func modularExponentiation(g, p, limit int) int {
	c := 1

	for i := 0; i < limit; i++ {
		c = (c * g) % p
	}

	return c
}
