package main

import (
	"fmt"

	"github.com/gookit/color"
)

// go run ./_examples/color_256.go
// FORCE_COLOR=on go run ./_examples/color_256.go
func main() {
	// var s *color.Style256

	fmt.Printf("%-45s 256 Color(16 bit) Table %-35s\n", " ", " ")
	// 0 - 16
	fmt.Printf("%-22sStandard Color %-42sExtended Color \n", " ", " ")
	for i := range []int{7: 0} { // 0 -7 -> 8/16color: 30–37
		color.S256(255, uint8(i)).Printf("   %-4d", i)
	}
	fmt.Print("    ")
	for i := range []int{7: 0} { // 8 -15 -> 16color: 90–97
		i += 8
		color.S256(0, uint8(i)).Printf("   %-4d", i)
	}

	var fg uint8 = 255
	fmt.Printf("\n%-50s216 Color\n", " ")
	for i := range []int{215: 0} { // 16-231：6 × 6 × 6 立方（216色）: 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
		v := i + 16

		if i != 0 {
			if i%18 == 0 {
				fg = 0
				fmt.Println() // new line
			}

			if i%36 == 0 {
				fg = 255
				// fmt.Println() // new line
			}
		}

		color.S256(fg, uint8(v)).Printf("  %-4d", v)
	}

	fmt.Printf("\n%-50s24th Order Grayscale Color\n", " ")
	for i := range []int{23: 0} { // // 232-255：从黑到白的24阶灰度色
		if i < 12 {
			fg = 255
		} else {
			fg = 0
		}

		i += 232
		color.S256(fg, uint8(i)).Printf(" %-4d", i)
	}
	fmt.Println()
}
