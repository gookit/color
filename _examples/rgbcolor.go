package main

import "github.com/gookit/color"

// go run ./_examples/rgbcolor.go
func main() {
	color.RGB(30, 144, 255).Println("message. use RGB number")

	color.HEX("#1976D2").Println("blue-darken")
	color.HEX("#D50000", true).Println("red-accent. use HEX style")

	color.RGBStyleFromString("213,0,0").Println("red-accent. use RGB number")
	color.HEXStyle("eee", "D50000").Println("deep-purple color")
}
