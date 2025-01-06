package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you haven't caught any pokemon")
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
