package main

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// go run ./_examples/colortag.go
func main() {
	i := 0
	fmt.Print("All color tags:\n\n")

	for tag := range color.GetColorTags() {
		if strings.Contains(tag, "_") {
			continue
		}

		i++
		color.Tag(tag).Printf("%-15s", tag)
		if i%5 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}

	fmt.Printf("\n\ntotal tags: %d\n", i)
}
