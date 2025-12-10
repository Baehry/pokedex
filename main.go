package main

import (
    "fmt"
    "bufio"
    "os"
    "github.com/Baehry/pokedex/internal/pokecache"
    "time"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    initCommands()
    c := &config{
        previous: "",
        next: "",
        cac: pokecache.NewCache(500 * time.Millisecond),
        arguments: make([]string, 0),
        pokemons: make(map[string]struct{}),
    }
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        text := cleanInput(scanner.Text())
        if len(text) == 0 {
            fmt.Print("No command entered")
            continue
        }
        if len(text) > 1 {
            c.arguments = text[1:]
        }
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