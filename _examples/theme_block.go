package main

import (
	"fmt"

	"github.com/gookit/color"
)

// go run _examples/theme_block.go
func main() {
	fmt.Println("Built In Themes(styles):")
	fmt.Println("\n------------------ BLOCK STYLE ------------------")

	for name, s := range color.Themes {
		s.Block("%s block style message", name)
	}
}
