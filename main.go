package main

import (
    "fmt"
    "strings"
)

func main() {
	fmt.Printf("Hello, World!")
}

func cleanInput(text string) []string {
    result := strings.Fields(text)
    for i, word := range result {
        result[i] = strings.ToLower(word)
    }
    return result
}

