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
            err := command.callback(c)
            if err != nil {
                fmt.Printf("%v\n", err)
                os.Exit(1)
            }
        } else {
            fmt.Print("Unknown command\n")
        }
    }
}