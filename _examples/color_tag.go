package main

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// go run ./_examples/color_tag.go
func main() {
	i := 0
	fmt.Println("Current env whether support color:", color.IsSupportColor())
	fmt.Print("All color tags:\n\n")

	for tag := range color.GetColorTags() {
		if strings.Contains(tag, "_") || strings.HasPrefix(tag, "hi") {
			continue
		}

		i++
		color.Tag(tag).Printf("tag: %-14s", tag)
		// taggedText := color.WrapTag("tag:" + tag, tag)
		// color.Printf("%s", taggedText)

		if i%5 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print("  ")
		}
	}

	fmt.Printf("\n\nBuilt-in Tags Total Number: %d\n", i)
}
