package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		cleanedTextSlice := cleanText(text)

		if len(cleanedTextSlice) == 0 {
			continue
		}

		commandName := cleanedTextSlice[0]

		availableCommands := getCommands()
		
		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		command.callback();

	}
}

func cleanText(text string) []string {
	lowerdText := strings.ToLower(text)
	textSlice := strings.Fields(strings.Trim(lowerdText, " "))
	return textSlice
}
