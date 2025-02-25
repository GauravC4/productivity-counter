package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"start": {
			name:        "start",
			description: "Start the clock.",
			callback:    commandStart,
		},
		"compress": {
			name:        "compress",
			description: "Compress WAL file by removing redundant sessions.",
			callback:    commandCompress,
		},
		"help": {
			name:        "help",
			description: "Display a help message.",
			callback:    commandHelp,
		},
		"analyse": {
			name:        "analyse",
			description: "Accepts 2 arguments <n> and <v>. Displays analysis for last n days with sessions if v.",
			callback:    commandAnalyse,
		},
		"exit": {
			name:        "exit",
			description: "Exit the productivity clock.",
			callback:    commandExit,
		},
	}
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Counter > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			return
		}

		cleanUserInputArr := CleanInput(scanner.Text())
		if len(cleanUserInputArr) == 0 {
			continue
		}

		command := cleanUserInputArr[0]
		args := cleanUserInputArr[1:]
		if cmd, ok := GetCommands()[command]; ok {
			err := cmd.callback(args)
			if err != nil {
				fmt.Printf("Error : %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
