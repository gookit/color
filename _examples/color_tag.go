package main

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// go run ./_examples/color_tag1.go
func main() {
	i := 0
	fmt.Println("Current env whether support color:", color.IsSupportColor())
	fmt.Print("All color tags:\n\n")

	for tag := range color.GetColorTags() {
		if strings.Contains(tag, "_") {
			continue
		}

		i++
		color.Tag(tag).Print(tag+" tag")
		if i%5 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print("  ")
		}
	}

	fmt.Printf("\n\ntotal tags: %d\n", i)
}
