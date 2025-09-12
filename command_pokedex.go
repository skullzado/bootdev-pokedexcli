package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		return errors.New("you have not caught any pokemon")
	}

	fmt.Println("Your Pokedex:")
	for k := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}

	return nil
}
