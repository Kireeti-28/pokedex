package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleanedTextSlice := cleanText(text)

		if len(cleanedTextSlice) == 0 {
			continue
		}

		args := []string{}
		if len(cleanedTextSlice) > 1 {
			args = cleanedTextSlice[1:]
		}

		commandName := cleanedTextSlice[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Print(err)
		}

	}
}

func cleanText(text string) []string {
	lowerdText := strings.ToLower(text)
	textSlice := strings.Fields(strings.Trim(lowerdText, " "))
	return textSlice
}
