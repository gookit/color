package main

import (
	"fmt"

	"github.com/gookit/color"
)

// go run ./_examples/color_16.go
func main() {
	fmt.Println("Foreground colors:")
	for name, c := range color.FgColors {
		c.Print(" ", name, " ")
	}

	fmt.Println("\nBackground colors:")
	for name, c := range color.BgColors {
		c.Print(" ", name, " ")
	}

	fmt.Println("\nBasic Options:")
	for name, c := range color.AllOptions {
		c.Print(" ", name, " ")
	}

	fmt.Println("\nExtra foreground colors:")
	for name, c := range color.ExFgColors {
		c.Print(" ", name, " ")
	}

	fmt.Println("\nExtra background colors:")
	for name, c := range color.ExBgColors {
		c.Print(" ", name, " ")
	}

	fmt.Println()
}
