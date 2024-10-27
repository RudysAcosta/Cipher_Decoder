package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

// Función de cifrado César que cifra o descifra un texto con un desplazamiento
func caesarCipher(text string, shift int) string {
	var result strings.Builder
	shift = (shift%26 + 26) % 26 // Asegura que el desplazamiento sea positivo entre 0 y 25
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			result.WriteRune('a' + (char-'a'+rune(shift))%26)
		} else if char >= 'A' && char <= 'Z' {
			result.WriteRune('A' + (char-'A'+rune(shift))%26)
		} else {
			result.WriteRune(char) // Deja caracteres especiales sin cambio
		}
	}
	return result.String()
}

func main() {
	var g, p, A int
	b := 7 // Valor secreto de Bob según el ejemplo

	// Paso 1: Leer g y p, y mostrar "OK"
	readInput("g is %d and p is %d", &g, &p)
	fmt.Println("OK")

	// Paso 2: Leer A y calcular B y el secreto compartido S
	readInput("A is %d", &A)
	B := modularExponentiation(g, p, b)
	S := modularExponentiation(A, p, b) // Secreto compartido

	// Imprimir B como se especifica
	fmt.Printf("B is %d\n", B)

	// Paso 3: Usar el secreto compartido S como desplazamiento para el cifrado César
	secretShift := S % 26 // Ajuste del desplazamiento para el cifrado

	// Cifrar el mensaje "Will you marry me?" usando el cifrado César
	message := "Will you marry me?"
	encryptedMessage := caesarCipher(message, secretShift)
	fmt.Println(encryptedMessage)

	// Paso 4: Leer la respuesta de Alice, descifrar y responder adecuadamente
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		response := scanner.Text()
		decodedResponse := caesarCipher(response, -secretShift) // Usar -shift para descifrar

		// Evaluar y responder en función del mensaje de Alice
		switch decodedResponse {
		case "Yeah, okay!":
			reply := caesarCipher("Great!", secretShift)
			fmt.Println(reply)
		case "Let's be friends.":
			reply := caesarCipher("What a pity!", secretShift)
			fmt.Println(reply)
		}
	}
}
