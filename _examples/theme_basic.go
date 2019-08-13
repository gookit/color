package main

import (
	"fmt"

	"github.com/gookit/color"
)

// go run _examples/theme_basic.go
func main() {
	fmt.Println("Built In Themes(styles):")
	fmt.Println("------------------ BASIC STYLE ------------------")

	for name, s := range color.Themes {
		s.Println(name, "message")
	}
}
