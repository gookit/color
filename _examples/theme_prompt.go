package main

import (
	"fmt"

	"github.com/gookit/color"
)

// go run _examples/theme_prompt.go
func main() {
	fmt.Println("Built In Themes(styles):")
	fmt.Println("\n------------------ PROMPT STYLE ------------------")

	for name, s := range color.Themes {
		s.Prompt("%s prompt message", name)
	}
}
