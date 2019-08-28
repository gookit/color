package main

import (
	"fmt"

	"github.com/gookit/color"
)

// go run _examples/theme_tips.go
func main() {
	fmt.Println("Built In Themes(styles):")
	fmt.Println("\n------------------ TIPS STYLE ------------------")

	for name, s := range color.Themes {
		s.Tips("%s tips message", name)
	}

	color.Info.Render()
}
