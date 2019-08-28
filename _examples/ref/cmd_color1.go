package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

// go run _examples/ref/cmd_color1.go
func main() {
	color.ForceOpenColor()

	ss := os.Environ()
	for _, v := range ss {
		fmt.Println(v)
	}

	str := color.Question.Render("msg", "More")
	fmt.Println(str)
}
