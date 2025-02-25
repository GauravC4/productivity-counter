package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Clock... Goodbye!")
	os.Exit(0)
	return nil
}
