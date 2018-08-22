package main

import (
	"fmt"
	"github.com/gookit/color"
)

// go run ./_examples/theme_style.go
func main()  {
	fmt.Println("Built in themes(styles):")
	fmt.Println("------------------ basic style ------------------")
	for name, s := range color.Themes {
		s.Print(" ", name, " ")
	}

	fmt.Println("\n------------------ tips style ------------------")
	for name, s := range color.Themes {
		s.Tips("%s tips message", name)
	}

	fmt.Println("------------------ prompt style ------------------")
	for name, s := range color.Themes {
		s.Prompt("%s prompt message", name)
	}
}
