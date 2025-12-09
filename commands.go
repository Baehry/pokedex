package main

import (
	"os"
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"time"
	"github.com/Baehry/pokedex/internal/pokecache"
)

var supportedCommands map[string]cliCommand
var cac pokecache.Cache

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	next string
	previous string
}

type locationarea struct {
	Count int
	Next string
	Previous string
	Results []struct{
		Name string
		Url string
	}
}

func commandExit(c *config) error {
    fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, value := range supportedCommands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.next != "" {
		url = c.next
	}
	data, ok := cac.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cac.Add(url, data)
	}
	var results locationarea
	if err := json.Unmarshal(data, &results); err != nil {
		return err
	}
	c.previous = results.Previous
	c.next = results.Next
	fmt.Print("Pokedex > map\n")
	for _, place := range results.Results {
		fmt.Printf("%s\n", place.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.previous == "" {
		fmt.Print("you're on the first page\n")
		return nil
	}
	data, ok := cac.Get(c.previous)
	if !ok {
		res, err := http.Get(c.previous)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cac.Add(c.previous, data)
	}
	var results locationarea
	if err := json.Unmarshal(data, &results); err != nil {
		return err
	}
	c.previous = results.Previous
	c.next = results.Next
	for _, place := range results.Results {
		fmt.Printf("%s\n", place.Name)
	}
	return nil
}

func initCommands() {
	supportedCommands = map[string]cliCommand {
    	"exit": {
        	name: "exit",
        	description: "Exit the Pokedex",
        	callback: commandExit,
    	},
		"help": {
        	name: "help",
        	description: "Displays a help message",
        	callback: commandHelp,
    	},
		"map": {
			name: "map",
			description: "Displays 1 page of the list of locations, each subsequent call displays the next page",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous page of the list displayed by the last map call",
			callback: commandMapb,
		},
	}
	cac = pokecache.NewCache(500 * time.Millisecond)
}
