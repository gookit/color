package main

import (
	"fmt"
	"time"

	"github.com/gookit/color"
)

func main() {
	green := color.HEX("#00FF62").Sprintf("Hello world!")
	now := time.Now().Format("02-01-2006 15:04:05.000")

	fmt.Printf("[%v] %v\n", now, green)

	color.Printf("[%v] %v\n", now, green)
	color.HEX("#00FF62").Println("Hello world!")
}
