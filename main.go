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

	fmt.Println(b)

	// B := int(math.Pow(float64(g), float64(b))) % p

	B := modularExponentiation(g, p, b)

	fmt.Printf("\nB is %d\n", B)
	fmt.Printf("A is %d", A)

	fmt.Println(a)

}

func modularExponentiation(g, p, limit int) int {
	c := 1

	for i := 0; i < limit; i++ {
		c = (c * g) % p
	}

	return c
}
