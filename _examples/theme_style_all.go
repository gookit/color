package main

import (
	"fmt"

	"github.com/gookit/color"
)

// go run ./_examples/theme_style_all.go
func main() {
	fmt.Println("Built In Themes(styles):")
	fmt.Println("------------------ Basic style ------------------")
	for name, s := range color.Themes {
		s.Println(name, "message ")
	}

	fmt.Println("\n------------------ Tips style ------------------")
	for name, s := range color.Themes {
		s.Tips("%s tips message", name)
	}

	fmt.Println("------------------ Prompt style ------------------")
	for name, s := range color.Themes {
		s.Prompt("%s prompt message", name)
	}
}
