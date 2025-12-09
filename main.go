package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    initCommands()
    c := &config{
        previous: "",
        next: "",
    }
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        text := cleanInput(scanner.Text())
        command, ok := supportedCommands[text[0]]
        if ok {
            if command.callback(c) != nil {
                os.Exit(1)
            }
        } else {
            fmt.Print("Unknown command\n")
        }
    }
}