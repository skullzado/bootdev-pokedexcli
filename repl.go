package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	re := regexp.MustCompile(`\P{L}+`)
	words := strings.Fields(output)
	for index, word := range words {
		w := re.ReplaceAllString(word, "")
		words[index] = w
	}
	return words
}
