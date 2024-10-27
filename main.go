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

func caesarCipher(text string, shift int) string {
	var result strings.Builder
	shift = (shift%26 + 26) % 26
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			result.WriteRune('a' + (char-'a'+rune(shift))%26)
		} else if char >= 'A' && char <= 'Z' {
			result.WriteRune('A' + (char-'A'+rune(shift))%26)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func main() {
	var g, p, A int
	b := 7

	readInput("g is %d and p is %d", &g, &p)
	fmt.Println("OK")

	readInput("A is %d", &A)
	B := modularExponentiation(g, p, b)
	S := modularExponentiation(A, p, b)

	fmt.Printf("B is %d\n", B)

	secretShift := S % 26

	message := "Will you marry me?"
	encryptedMessage := caesarCipher(message, secretShift)
	fmt.Println(encryptedMessage)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		response := scanner.Text()
		decodedResponse := caesarCipher(response, -secretShift)

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
