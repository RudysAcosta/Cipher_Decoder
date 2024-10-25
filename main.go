package main

import (
	"fmt"
	"log"
)

func readInput(format string, vars ...interface{}) {
	_, err := fmt.Scanf(format, vars...)
	if err != nil {
		log.Fatalf("Error al leer los valores: %v", err)
	}
}

// Exponenciación modular por cuadrados
func modularExponentiation(base, mod, exp int) int {
	result := 1
	base = base % mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp = exp / 2
	}
	return result
}

func main() {
	var g, p, A int
	a, b := 21, 15

	readInput("g is %d and p is %d", &g, &p)
	fmt.Println("OK")
	readInput("A is %d", &A)

	B := modularExponentiation(g, p, b)
	Sa := modularExponentiation(A, p, b)
	Sb := modularExponentiation(B, p, a)

	fmt.Printf("B is %d\n", B)
	fmt.Printf("A is %d\n", A)

	if Sa == Sb {
		fmt.Printf("S is %d\n", Sa)
	} else {
		log.Fatal("Algo salió mal: Sa y Sb no coinciden")
	}
}
