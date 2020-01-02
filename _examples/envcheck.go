package main

import (
	"fmt"

	"github.com/gookit/color"
)

func main() {
	fmt.Println("IsSupport256Color", color.IsSupport256Color())
	fmt.Println("IsSupportColor", color.IsSupportColor())
}
