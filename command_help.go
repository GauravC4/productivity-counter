package main

import "fmt"

func commandHelp(args []string) error {
	fmt.Println("Welcome to the Productivity Clock!")
	fmt.Print("\nUsage:\n\n")
	for _, val := range GetCommands() {
		fmt.Printf("%v: %v\n\n", val.name, val.description)
	}
	return nil
}
