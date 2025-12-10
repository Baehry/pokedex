package main

import (
	"os"
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"github.com/Baehry/pokedex/internal/pokecache"
	"math/rand"
)

var supportedCommands map[string]cliCommand

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	next string
	previous string
	cac pokecache.Cache
	arguments []string
	pokemons map[string]struct{}
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
	data, ok := c.cac.Get(url)
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
		c.cac.Add(url, data)
	}
	var results NamedAPIResourceList
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
	data, ok := c.cac.Get(c.previous)
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
		c.cac.Add(c.previous, data)
	}
	var results NamedAPIResourceList
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

func commandExplore(c *config) error {
	if len(c.arguments) < 1 {
		fmt.Print("No Argument Given")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + c.arguments[0] + "/"
	data, ok := c.cac.Get(url)
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
		c.cac.Add(url, data)
	}
	var result LocationArea
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	for _, encounter := range result.PokemonEncounters {
		fmt.Printf("%s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(c *config) error {
	if len(c.arguments) < 1 {
		fmt.Print("No Argument Given")
		return nil
	}
	if _, ok := c.pokemons[c.arguments[0]]; ok {
		fmt.Print("Pokemon was already caught\n")
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon-species/" + c.arguments[0] + "/"
	data, ok := c.cac.Get(url)
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
		c.cac.Add(url, data)
	}
	var result PokemonSpecies
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", result.Name)
	if rand.Intn(256) <= result.CaptureRate {
		fmt.Printf("%s was caught!\n", result.Name)
		c.pokemons[result.Name] = struct{}{}
		return nil
	}
	fmt.Printf("%s escaped!\n", result.Name)
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
		"explore": {
			name: "explore",
			description: "Displays a list of wild Pokemon available in a specified area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempts to catch a wild Pokemon",
			callback: commandCatch,
		},
	}
}
