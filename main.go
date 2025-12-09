package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        text := cleanInput(scanner.Text())
        if len(text) == 0 {
            return
        }
        fmt.Printf("Your command was: %s\n", text[0])
    }
}