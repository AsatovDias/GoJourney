package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Конкатенация
	first := "Go"
	second := "Lang"
	full := first + " " + second

	fmt.Println("Конкатенация:", full)

	// 2. Длина строки (в байтах)
	s := "Go Язык"
	fmt.Println("len:", len(s))

	// 3. strings.Contains
	fmt.Println("Contains 'Go':", strings.Contains(s, "Go"))

	// 4. strings.Split
	words := strings.Split(s, " ")
	fmt.Println("Split:", words)

	// 5. strings.Replace
	replaced := strings.Replace(s, "Go", "Golang", 1)
	fmt.Println("Replace:", replaced)

	// 6. byte vs rune
	fmt.Println("\n--- byte (по байтам) ---")
	for i := 0; i < len(s); i++ {
		fmt.Printf("index: %d, byte: %v\n", i, s[i])
	}

	fmt.Println("\n--- rune (по символам) ---")
	for i, r := range s {
		fmt.Printf("index: %d, rune: %c, code: %d\n", i, r, r)
	}
}
